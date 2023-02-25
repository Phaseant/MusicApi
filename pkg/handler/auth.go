package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
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

func (h *Handler) signIn(c *gin.Context) {

}
