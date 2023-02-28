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

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) //sign up
		auth.POST("/sign-in", h.signIn) //sign in
	}

	api := router.Group("/api", h.userIdentity)
	{
		album := api.Group("/album")
		{
			album.POST("/", h.createAlbum) //create new album
			album.GET("/", h.getAllAlbums) //get all albums

			album.GET("/:id", h.getAlbumById)       //get album by its id
			album.DELETE("/:id", h.deleteAlbumById) //delete album
		}
	}
	return router
}
