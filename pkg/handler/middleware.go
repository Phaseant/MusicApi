package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) userIdentity(c *gin.Context) {
	token := strings.Split(c.Request.Header["Authorization"][0], " ")[1]

	if len(token) == 0 {
		log.Errorf("Error: empty auth header")
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Empty auth header"})
		return
	}

	userID, err := h.services.Autorization.ParseToken(token)
	if err != nil {
		log.Errorf("Error parsing token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Not valid token"})
		return
	}

	c.Set(userctx, userID)
}

func getUserID(c *gin.Context) (string, error) {
	id := c.GetString(userctx)
	if len(id) == 0 {
		return "", fmt.Errorf("user id %v not found", id)
	}

	return id, nil
}
