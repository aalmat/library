package repository

import (
	"github.com/aalmat/bookstore/models"
	"github.com/jinzhu/gorm"
)

type BookPostgres struct {
	db *gorm.DB
}

func NewBookPostgres(db *gorm.DB) *BookPostgres {
	return &BookPostgres{db}
}

func (b *BookPostgres) GetBooks() ([]models.Book, error) {
	var books []models.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
