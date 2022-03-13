package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetAllBooks(c *gin.Context) {

	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	var booksResponse []book.BookResponse
	for _, b := range books {
		bookResponse := convertToBookResponse(b)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse,
	})
}

func (h *bookHandler) GetSingleBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	bookResponse := convertToBookResponse(b)

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func convertToBookResponse(b book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}

// func (h *bookHandler) RootHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": "Faris Fahmi",
// 		"bio":  "Backend Engineer",
// 	})
// }

// func (h *bookHandler) QueryHandler(c *gin.Context) {
// 	title := c.Query("title")
// 	price := c.Query("price")

// 	c.JSON(http.StatusOK, gin.H{
// 		"title": title,
// 		"price": price,
// 	})
// }

// func (h *bookHandler) BooksHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	title := c.Param("title")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id":    id,
// 		"title": title,
// 	})
// }

// func (h *bookHandler) HelloHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"title":    "Hello World",
// 		"subtitle": "Belajar Golang Backend Engineer",
// 	})
// }
