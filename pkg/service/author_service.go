package service

import (
	"github.com/aalmat/bookstore/models"
	"github.com/aalmat/bookstore/pkg/repository"
)

type AuthorService struct {
	repo repository.Author
}

func NewAuthorService(repo repository.Author) *AuthorService {
	return &AuthorService{
		repo,
	}
}

func (b *AuthorService) Create(authorId uint, book models.Book) (uint, error) {
	return b.repo.Create(authorId, book)
}

func (b *AuthorService) GetBook(bookId uint) (models.Book, error) {
	return b.repo.GetBook(bookId)
}

func (b *AuthorService) Update(bookId uint, update models.UpdateBook) error {
	return b.repo.Update(bookId, update)
}

func (b *AuthorService) Delete(id uint) error {
	return b.repo.Delete(id)
}

func (b *AuthorService) GetAuthorBooks(id uint) ([]models.Book, error) {
	return b.repo.GetAuthorBooks(id)
}
