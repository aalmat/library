package handler

import (
	"github.com/aalmat/bookstore/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		s,
	}
}

func (h *Handler) Routes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
	}

	author := router.Group("/author", h.UserIdentify)
	{
		author.POST("/", h.Create)
		author.GET("/", h.GetAuthorBooks)
		author.GET("/:id", h.GetBook)
		author.PUT("/:id", h.Update)
		author.DELETE("/:id", h.Delete)
	}

	member := router.Group("/member", h.UserIdentify)
	{
		member.GET("/", h.GetClientBooks)
		member.POST("/:id", h.OrderBook)
		member.DELETE("/:id", h.DeleteOrder)
	}

	book := router.Group("/book")
	{
		book.GET("/", h.GetBooks)
	}

	return router
}
