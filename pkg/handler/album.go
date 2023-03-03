package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) createAlbum(c *gin.Context) {
	_, err := getUserID(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	var album entity.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		log.Errorf("CreateAlbum: Error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	id, err := h.services.Album.NewAlbum(album)

	if mongo.IsDuplicateKeyError(err) { //key values : title and author together
		log.Errorf("%v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "this album is already exists"})
		return
	}

	if err != nil {
		log.Errorf("Error trying to add new album: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ID": id})
}

func (h *Handler) getAllAlbums(c *gin.Context) {
	albums, err := h.services.GetAllAlbums()
	if err != nil {
		log.Errorf("Error getting all albums: %v")
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (h *Handler) getAlbum(c *gin.Context) {
	albumId := c.Param("albumID")

	album, err := h.services.GetAlbum(albumId)

	if err != nil {
		log.Errorf("Error finding album: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *Handler) deleteAlbum(c *gin.Context) {
	albumId := c.Param(albumctx)

	_, err := getUserID(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ok := h.services.DeleteAlbum(albumId)
	if !ok {
		log.Error("error deleting album")
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "album was not deleted"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Deleted": true})
}
