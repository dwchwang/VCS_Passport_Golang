package handler

import (
	"net/http"
	"strconv"

	"github.com/dwchwang/rest_api_golang/internal/service"
	"github.com/gin-gonic/gin"
)

type BorrowRecordHandler struct {
	service service.BorrowRecordService
}

func NewBorrowRecordHandler(service service.BorrowRecordService) *BorrowRecordHandler {
	return &BorrowRecordHandler{service: service}
}

type BorrowBookRequest struct {
	BorrowerID uint `json:"borrower_id" binding:"required"`
	BookID     uint `json:"book_id" binding:"required"`
}

func (h *BorrowRecordHandler) BorrowBook(c *gin.Context) {
	var req BorrowBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.BorrowBook(req.BorrowerID, req.BookID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
}


func (h *BorrowRecordHandler) ReturnBook(c *gin.Context) {
	recordIDParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid record ID"})
		return
	}
	if err := h.service.ReturnBook(uint(recordIDParam)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}
