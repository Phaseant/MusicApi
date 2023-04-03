package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) addAdmin(c *gin.Context) {
	ok := adminFlag(c)
	if !ok {
		return
	}

	var newAdmin entity.Admin

	if err := c.ShouldBindJSON(&newAdmin); err != nil {
		log.Errorf("addAdmin: error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"added": false, "error": err.Error()})
		return
	}

	err := h.services.Admin.AddAdmin(newAdmin)
	if err != nil {
		log.Errorf("addAdmin: unable to add new admin: %v", err)
	}

	if mongo.IsDuplicateKeyError(err) {
		log.Errorf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"added": false, "error": "unable to add new admin, this user is already an admin"})
		return
	}

	if err != nil {
		log.Errorf("addAdmin: error trying to add new admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"added": false, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"added": true})
}
