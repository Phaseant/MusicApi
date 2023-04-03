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
		log.Errorf("admintIdentity: no authorization token provided")
		c.JSON(http.StatusBadRequest, gin.H{"error": "authorization token not provided"})
		return
	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		log.Errorf("adminIdentity: could not find bearer token in Authorization header")
		c.JSON(http.StatusForbidden, gin.H{"error": "unable to find bearer token in Authorization header"})
		return
	}

	userID, err := h.services.Autorization.ParseToken(token)
	if err != nil {
		log.Errorf("adminIdentity: error parsing token: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "provided token is invalid"})
		return
	}
	if h.services.Admin.IsAdmin(userID) {
		log.Info("adminIdentity: admin with id %s sent a request", userID)
		c.Set(userctx, userID)
	}
}

func adminFlag(c *gin.Context) bool {
	id := c.GetString(userctx)
	return len(id) != 0
}
