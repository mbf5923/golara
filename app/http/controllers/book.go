package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm-test/app/http/requests"
	"gorm-test/app/http/resources"
	"gorm-test/app/models"
	"gorm-test/utils"
	"net/http"
)

type BookInterface interface {
	CreateBook(c *gin.Context)
}

type BookRepo struct {
	BookModel *models.BookModel
}

func NewBookRepo() *BookRepo {
	return &BookRepo{
		BookModel: models.NewBookModel(),
	}
}

func (repository *BookRepo) CreateBook(ctx *gin.Context) {
	var book models.Book
	if !requests.BookCreateRequestHandler(ctx, &book) {
		return
	}
	user, _ := ctx.Get("user")
	book.UserId = user.(*models.User).ID
	err := repository.BookModel.CreateBook(&book)
	if err != nil {
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	var bookResource resources.BookResource
	response := utils.Responses{}
	response.MakeResponse(book, &bookResource).SuccessResponse(ctx, bookResource, "ok", 200)
}

func (repository *BookRepo) Books(ctx *gin.Context) {
	var book []models.Book
	user, _ := ctx.Get("user")
	userId := user.(*models.User).ID
	err := repository.BookModel.Books(int(userId), &book)
	if err != nil {
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	var bookResource []resources.BookResource
	response := utils.Responses{}
	response.MakeResponse(book, &bookResource).SuccessResponse(ctx, bookResource, "ok", 200)
}
