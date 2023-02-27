package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	//TODO add validation if user is already exists
	var user entity.User
	if err := c.BindJSON(&user); err != nil {
		log.Errorf("Error while binding User JSON: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Autorization.NewUser(user)
	if err != nil {
		log.Errorf("error trying to create new user: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{"id": id})
}

type userSignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var user userSignIn
	if err := c.BindJSON(&user); err != nil {
		log.Errorf("Error while binding User JSON: %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := h.services.Autorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		log.Errorf("error trying get token: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusAccepted, map[string]interface{}{"token": token})

}
