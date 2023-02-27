package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) signUp(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("Error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := h.services.Autorization.NewUser(user)
	if mongo.IsDuplicateKeyError(err) {
		log.Errorf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "This user is already exists!"})
		return
	}
	if err != nil {
		log.Errorf("Error trying to create new user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("Error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Autorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		log.Errorf("Error trying to get token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token})

}
