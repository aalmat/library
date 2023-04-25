package repository

import (
	"github.com/aalmat/bookstore/models"
	"github.com/jinzhu/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error) // id, err
	GetUser(username, password string) (models.User, error)
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

type Repository struct {
	Authorization
	Book
	Author
	Member
}

func NewPostgres(db *gorm.DB) *Repository {
	return &Repository{
		NewAuthPostgres(db),
		NewBookPostgres(db),
		NewAuthorPostgres(db),
		NewMemberPostgres(db),
	}
}
