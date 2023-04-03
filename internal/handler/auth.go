package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) register(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("register: error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Autorization.NewUser(user)

	if mongo.IsDuplicateKeyError(err) {
		log.Errorf("register: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "this user already exists"})
		return
	}

	if err != nil {
		log.Errorf("register: error trying to create new user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error, try again later"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) login(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("login: error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Autorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		log.Errorf("login: error trying to create token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
