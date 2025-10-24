package controller

import (
	"cinema-app/internal/dto"
	"cinema-app/internal/model"
	"cinema-app/internal/service"
	"cinema-app/internal/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redsync/redsync/v4"
)

type TransactionController struct {
	service    *service.TransactionService
	seatSrv    service.SeatService
	seatTrxSrv service.SeatTrxSrv
}

func NewTransactionController(s *service.TransactionService, seat service.SeatService, seatTrx service.SeatTrxSrv) *TransactionController {
	return &TransactionController{
		service:    s,
		seatSrv:    seat,
		seatTrxSrv: seatTrx,
	}
}

type CreateTransactionRequest struct {
	UserID        string                   `json:"user_id"`
	ShowtimeID    string                   `json:"showtime_id"`
	Seat          []*model.SeatTransaction `json:"seats"`
	PaymentMethod string                   `json:"payment_method"`
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var mutexes []*redsync.Mutex
	for _, s := range req.Seat {
		mutex := utils.Rs.NewMutex(fmt.Sprintf("lock:seat:%s", s.SeatID))
		if err := mutex.Lock(); err != nil {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("seat %s is being ordered", s.SeatID),
			})
			// Unlock all that locked before
			for _, m := range mutexes {
				_, _ = m.Unlock()
			}
			return
		}
		mutexes = append(mutexes, mutex)
	}

	//check seat
	for _, s := range req.Seat {
		seat, err := c.seatSrv.GetByID(s.SeatID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "seat not found"})
			return
		}
		if seat.Status == "booking" || seat.Status == "paid" {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": fmt.Sprintf("seat %s is already booked", seat.SeatNumber),
			})
			return
		}
	}

	t, err := c.service.CreateTransaction(req.UserID, req.ShowtimeID, req.Seat, req.PaymentMethod)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "transaction created", "data": t})

	//update seat model
	var filter dto.Seat
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids := []string{}
	for _, id := range req.Seat {
		ids = append(ids, id.SeatID)
	}
	if len(ids) > 0 {
		filter.IDs = ids
	}

	updates := model.Seat{
		Status:   "booking",
		IsBooked: true,
	}
	updatesSeatTrx := model.SeatTransaction{
		Status: "booking",
		IsPaid: false,
	}

	if err := c.seatSrv.UpdateSeats(filter, t.ID, &updates, &updatesSeatTrx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// unlock mutex
	for _, m := range mutexes {
		_, _ = m.Unlock()
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seat updated"})
}

func (c *TransactionController) MarkAsPaid(ctx *gin.Context) {
	idParam := ctx.Param("id")

	err := c.service.MarkAsPaid(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction marked as paid"})

	//update seat
	var filter dto.SeatTrx
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter.TransactionID = idParam
	strx, err := c.seatTrxSrv.GetsSeatTrx(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filterSeat dto.Seat
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids := []string{}
	for _, id := range *strx {
		ids = append(ids, id.SeatID)
	}
	if len(ids) > 0 {
		filterSeat.IDs = ids
	}

	updates := model.Seat{
		Status:   "paid",
		IsBooked: true,
	}
	updatesSeatTrx := model.SeatTransaction{
		Status: "paid",
		IsPaid: true,
	}

	if err := c.seatSrv.UpdateSeats(filterSeat, idParam, &updates, &updatesSeatTrx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Seat updated"})

}

func (c *TransactionController) CancelOrder(ctx *gin.Context) {
	idParam := ctx.Param("id")
	err := c.service.CancelOrder(idParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Transaction marked as canceled"})

	//update seat
	var filter dto.SeatTrx
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filter.TransactionID = idParam
	strx, err := c.seatTrxSrv.GetsSeatTrx(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var filterSeat dto.Seat
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids := []string{}
	for _, id := range *strx {
		ids = append(ids, id.SeatID)
	}
	if len(ids) > 0 {
		filterSeat.IDs = ids
	}

	updates := model.Seat{
		Status:   "available",
		IsBooked: false,
	}
	updatesSeatTrx := model.SeatTransaction{
		Status: "canceled",
		IsPaid: false,
	}

	if err := c.seatSrv.UpdateSeats(filterSeat, idParam, &updates, &updatesSeatTrx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Seat updated"})

}
