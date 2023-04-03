package handler

import (
	"github.com/Phaseant/MusicAPI/internal/service"
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
			album.POST("/", h.adminIdentity, h.addAlbum)              //add new album
			album.POST("/array", h.adminIdentity, h.addAlbums)        //add multiple albums [array]
			album.DELETE("/:albumID", h.adminIdentity, h.deleteAlbum) //delete album
		}

		admin := api.Group("/admin")
		{
			admin.POST("/", h.adminIdentity, h.addAdmin) //add new admin
		}
	}

	return router
}
