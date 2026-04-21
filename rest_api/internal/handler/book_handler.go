package handler

import (
	"net/http"
	"strconv"

	"github.com/dwchwang/rest_api_golang/internal/models"
	"github.com/dwchwang/rest_api_golang/internal/service"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

// POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Thêm sách thành công", "data": book})
}

// GET /books
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GET /books/:id
func (h *BookHandler) GetBookById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là một số nguyên"})
		return
	}

	book, err := h.service.GetBookById(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// PUT /books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	var updatedData models.Book
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	if err := h.service.UpdateBook(uint(id), &updatedData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật sách thành công"})
}

// DELETE /books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	if err := h.service.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa sách thành công"})
}

func (h *BookHandler) GetUnborrowedBooks(c *gin.Context) {
	books, err := h.service.GetUnborrowedBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi lấy danh sách sách ế: " + err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Lấy danh sách thành công",
		"total":   len(books), 
		"data":    books,
	})
}