package service

import (
	"github.com/aalmat/bookstore/models"
	"github.com/aalmat/bookstore/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error) // id, err
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, models.Role, error) // id err
}

type Book interface {
	GetBooks() ([]models.Book, error)
}

type Author interface {
	Create(authorId uint, book models.Book) (uint, error)
	GetBook(bookId uint) (models.Book, error)
	Update(bookId uint, update models.UpdateBook) error
	Delete(id uint) error
	GetAuthorBooks(id uint) ([]models.Book, error)
}

type Member interface {
	GetClientBooks(id uint) ([]models.Book, error)
	OrderBook(userId, bookId uint) (uint, error)
	DeleteOrder(userId, bookId uint) error
}

type Service struct {
	Authorization
	Book
	Author
	Member
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewAuthService(repo),
		NewBookService(repo),
		NewAuthorService(repo),
		NewMemberService(repo),
	}
}
