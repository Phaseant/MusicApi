package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Errorf("No authorization token provided")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No authorization token provided"})
		return
	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		log.Errorf("Could not find bearer token in Authorization header")
		c.JSON(http.StatusForbidden, gin.H{"Error": "Could not find bearer token in Authorization header"})
		return
	}

	userID, err := h.services.Autorization.ParseToken(token)
	if err != nil {
		log.Errorf("Error parsing token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Not valid token"})
		return
	}
	if h.services.Admin.IsAdmin(userID) {
		log.Info("Admin found. id: ", userID)
		c.Set(userctx, userID)
	}
}

func adminFlag(c *gin.Context) bool {
	id := c.GetString(userctx)
	return len(id) != 0
}
