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

	// Lễ tân nhận JSON từ Postman và map nó vào biến 'book'
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	// Gọi Quản lý (Service) để xử lý
	if err := h.service.CreateBook(&book); err != nil {
		// Ném lỗi của Service ra cho Client (ví dụ lỗi thiếu tên sách)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Trả về thành công
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
	// Lấy ID từ URL (vd: /books/1) và ép kiểu từ chuỗi sang số nguyên
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là một số nguyên"})
		return
	}

	book, err := h.service.GetBookById(uint(id))
	if err != nil {
		// Trả về mã 404 (Not Found) nếu không tìm thấy
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
