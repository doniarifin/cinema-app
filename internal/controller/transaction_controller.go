package controller

import (
	"cinema-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *service.TransactionService
}

func NewTransactionController(s *service.TransactionService) *TransactionController {
	return &TransactionController{s}
}

type CreateTransactionRequest struct {
	UserID        string `json:"user_id"`
	ShowtimeID    string `json:"showtime_id"`
	SeatID        string `json:"seat_id"`
	PaymentMethod string `json:"payment_method"`
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := c.service.CreateTransaction(req.UserID, req.ShowtimeID, req.SeatID, req.PaymentMethod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "transaction created", "data": t})
}

func (c *TransactionController) MarkAsPaid(ctx *gin.Context) {
	idParam := ctx.Param("id")

	err := c.service.MarkAsPaid(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction marked as paid"})
}
