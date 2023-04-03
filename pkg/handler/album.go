package handler

import (
	"net/http"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

func (h *Handler) createAlbum(c *gin.Context) {
	ok := adminFlag(c)
	if !ok {
		return
	}

	var album entity.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		log.Errorf("createAlbum: error while binding User JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.services.Album.NewAlbum(album)

	if mongo.IsDuplicateKeyError(err) {
		log.Errorf("createAlbum: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to add album, this album already exists"})
		return
	}

	if err != nil {
		log.Errorf("createAlbum: error adding new album: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"added": true, "id": id})
}

func (h *Handler) getAllAlbums(c *gin.Context) {
	albums, err := h.services.GetAllAlbums()
	if err != nil {
		log.Errorf("getAllAlbums: error getting all albums: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, albums)
}

func (h *Handler) getAlbum(c *gin.Context) {
	albumId := c.Param("albumID")

	album, err := h.services.GetAlbum(albumId)

	if err != nil {
		log.Errorf("getAlbum: error finding album with id %s: %v", albumId, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

func (h *Handler) deleteAlbum(c *gin.Context) {
	albumId := c.Param(albumctx)

	ok := adminFlag(c)
	if !ok {
		return
	}

	ok = h.services.DeleteAlbum(albumId)
	if !ok {
		log.Info("album with id %s not found: ", albumId)
		c.JSON(http.StatusBadRequest, gin.H{"deleted": false, "error": "album not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Deleted": true})
}
