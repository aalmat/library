package handler

import (
	"github.com/aalmat/bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetBook(ctx *gin.Context) {
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	book, err := h.service.Author.GetBook(uint(bookId))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (h *Handler) Update(ctx *gin.Context) {
	var update models.UpdateBook
	if err := ctx.BindJSON(&update); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Author.Update(uint(bookId), update)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": bookId})
}

func (h *Handler) Delete(ctx *gin.Context) {
	bookId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Author.Delete(uint(bookId))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": bookId})

}

func (h *Handler) Create(ctx *gin.Context) {

	authorId, err := h.GetUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	role, err := h.GetUserRole(ctx)

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if role != models.Author {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	var input models.Book
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Author.Create(authorId, input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) GetAuthorBooks(ctx *gin.Context) {
	authorId, err := h.GetUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	role, err := h.GetUserRole(ctx)

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if role != models.Author {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	books, err := h.service.Author.GetAuthorBooks(authorId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, BookList{
		books,
	})

}
