package handler

import (
	"github.com/aalmat/bookstore/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetClientBooks(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	role, err := h.GetUserRole(ctx)

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	if role != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	books, err := h.service.Member.GetClientBooks(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, BookList{
		books,
	})

}

func (h *Handler) OrderBook(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	userRole, err := h.GetUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	if userRole != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	bookId := ctx.Param("id")
	idUint, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Member.OrderBook(userId, uint(idUint))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) DeleteOrder(ctx *gin.Context) {
	userId, err := h.GetUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	userRole, err := h.GetUserRole(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	if userRole != models.Client {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	bookId := ctx.Param("id")
	idUint, err := strconv.ParseUint(bookId, 10, 64)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteOrder(userId, uint(idUint))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
