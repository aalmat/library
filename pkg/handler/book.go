package handler

import (
	"github.com/aalmat/bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookList struct {
	Data []models.Book
}

func (h *Handler) GetBooks(ctx *gin.Context) {
	books, err := h.service.Book.GetBooks()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, BookList{books})
}
