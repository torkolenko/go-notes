package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/torkolenko/go-notes/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	notes := router.Group("/notes")
	{
		notes.GET("/", h.getAllNotes)
		notes.GET("/:id", h.getNoteById)
		notes.POST("/", h.createNote)
		notes.PUT("/:id", h.updateNote)
		notes.DELETE("/:id", h.deleteNote)
	}

	return router
}
