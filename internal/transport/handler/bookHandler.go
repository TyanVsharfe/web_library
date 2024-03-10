package handler

import (
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type handlerBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
	Pages       int    `json:"pages"`
	Author      string `json:"author"`
}

func CreateBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book handlerBook

		// login := c.GetString("user")

		if err := c.BindJSON(&book); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		// book.Author = login

		id, err := service.CreateBook(c.Request.Context(), model.Book(book))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"book": id})

	}
}

func GetBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id книги"})

			return
		}

		book, err := service.GetBook(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения книги"})

			return

		}

		c.JSON(http.StatusOK, handlerBook(book))

	}
}

func GetBooks(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var books []handlerBook

		rBooks, err := service.GetBooks(c.Request.Context())

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения книги"})

			return

		}

		for _, book := range rBooks {
			books = append(books, handlerBook(book))
		}

		c.JSON(http.StatusOK, books)

	}
}

func DeleteBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id книги"})

			return
		}

		err = service.DeleteBook(c.Request.Context(), numberId)

		if err != nil {
			slog.Error(err.Error())

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "ошибка получения книги"})

			return

		}

		c.JSON(http.StatusOK, gin.H{"book deleted": id})

	}
}

func UpdateBook(service service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book handlerBook

		id := c.Param("id")

		numberId, err := strconv.Atoi(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверно передан id книги"})

			return
		}

		if err := c.BindJSON(&book); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		uBook, err := service.UpdateBook(c.Request.Context(), model.Book(book), numberId)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.JSON(http.StatusOK, handlerBook(uBook))

	}
}
