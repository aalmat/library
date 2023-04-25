package service

import (
	"github.com/aalmat/bookstore/models"
	"github.com/aalmat/bookstore/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo}
}

func (b *BookService) GetBooks() ([]models.Book, error) {
	return b.repo.GetBooks()
}
