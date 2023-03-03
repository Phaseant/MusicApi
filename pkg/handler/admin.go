package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) addAdmin(c *gin.Context) {
	_, err := getUserID(c)
	if err != nil {
		log.Error(err)
		return
	}

	var newAdmin entity.Admin

	if err := c.ShouldBindJSON(&newAdmin); err != nil {
		log.Errorf("addAdmin: Error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	err = h.services.Admin.AddAdmin(newAdmin)
	if err != nil {
		log.Error("Unable to add new admin: %v", err)
	}

	if mongo.IsDuplicateKeyError(err) {
		log.Errorf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "this id is already admin"})
		return
	}

	if err != nil {
		log.Errorf("Error trying to add new admin: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Added": true})
}
