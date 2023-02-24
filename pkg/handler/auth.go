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

	c.JSON(http.StatusAccepted, &user) //??

	h.services.Autorization.NewUser(user)
}

func (h *Handler) signIn(c *gin.Context) {

}
