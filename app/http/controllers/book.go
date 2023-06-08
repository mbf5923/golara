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

	err := repository.BookModel.CreateBook(&book)
	if err != nil {
		utils.FailedResponse(ctx, "Server Error", http.StatusInternalServerError, err)
		return
	}
	var bookResource resources.BookResource
	//utils.MakeResponse(book, &bookResource)
	response := utils.Responses{}
	response.MakeResponse(book, &bookResource).SuccessResponse(ctx, bookResource, "ok", 200)
	//utils.SuccessResponse(ctx, bookResource, "ok", 200)
}
