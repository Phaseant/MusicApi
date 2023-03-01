package handler

import (
	"github.com/Phaseant/MusicAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

const (
	userctx  = "userID"
	albumctx = "albumID"
)

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register) //sign up
		auth.POST("/login", h.login)       //sign in
	}

	api := router.Group("/api")
	{
		album := api.Group("/album")
		{
			album.GET("/", h.getAllAlbums)     //get all albums
			album.GET("/:albumID", h.getAlbum) //get album by its id

			//admin
			album.POST("/", h.userIdentity, h.createAlbum)           //create new album
			album.DELETE("/:albumID", h.userIdentity, h.deleteAlbum) //delete album
		}
	}
	return router
}
