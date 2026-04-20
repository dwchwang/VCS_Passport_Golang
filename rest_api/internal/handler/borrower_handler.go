package handler

import (
	"net/http"
	"strconv"

	"github.com/dwchwang/rest_api_golang/internal/models"
	"github.com/dwchwang/rest_api_golang/internal/service"
	"github.com/gin-gonic/gin"
)

type BorrowerHandler struct {
	service service.BorrowerService
}

func NewBorrowerHandler(s service.BorrowerService) *BorrowerHandler {
	return &BorrowerHandler{service: s}
}

func (h *BorrowerHandler) CreateBorrower(c *gin.Context) {
	var borrower models.Borrower
	if err := c.ShouldBindJSON(&borrower); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	if err := h.service.CreateBorrower(&borrower); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Thêm người mượn thành công", "data": borrower})
}

func (h *BorrowerHandler) GetAllBorrowers(c *gin.Context) {
	borrowers, err := h.service.GetAllBorrowers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": borrowers})
}

func (h *BorrowerHandler) GetBorrowerById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số nguyên"})
		return
	}

	borrower, err := h.service.GetBorrowerById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": borrower})
}

func (h *BorrowerHandler) UpdateBorrower(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var updatedData models.Borrower
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	if err := h.service.UpdateBorrower(uint(id), &updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật thành công"})
}

func (h *BorrowerHandler) DeleteBorrower(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	if err := h.service.DeleteBorrower(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
}