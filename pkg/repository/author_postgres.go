package repository

import (
	"github.com/aalmat/bookstore/models"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

type AuthorPostgres struct {
	db *gorm.DB
}

func NewAuthorPostgres(db *gorm.DB) *AuthorPostgres {
	return &AuthorPostgres{db}
}

func (b *AuthorPostgres) Create(authorId uint, book models.Book) (uint, error) {
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	book.UserId = authorId
	if err := b.db.Select("title", "genres", "created_at", "updated_at", "isbn", "user_id").Create(&book).Error; err != nil {
		return 0, err
	}

	return book.ID, nil
}

func (b *AuthorPostgres) GetBooksCostDesc() ([]models.Book, error) {
	var books []models.Book
	if err := b.db.Order("cost desc").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}
func (b *AuthorPostgres) Update(bookId uint, update models.UpdateBook) error {

	var book models.Book
	if err := b.db.First(&book, bookId).Error; err != nil {
		return err
	}
	if err := UpdateUtil(&book, update); err != nil {
		return err
	}

	err := b.db.Save(&book).Error
	if err != nil {
		return err
	}

	return nil
}
func UpdateUtil(b *models.Book, update models.UpdateBook) error {
	if err := update.Validate(); err != nil {
		return err
	}
	if update.Title != "" {
		b.Title = update.Title
	}
	if update.Title != "" {
		b.Isbn = update.Isbn
	}

	b.UpdatedAt = time.Now()
	return nil

}

func (b *AuthorPostgres) Delete(id uint) error {
	err := b.db.Where("id = ?", id).Delete(models.Book{}).Error
	return err
}

func (b *AuthorPostgres) GetAuthorBooks(id uint) ([]models.Book, error) {
	var books []models.Book

	if err := b.db.Where("user_id = ?", id).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func getBookId(bookIds []string) []uint {
	ids := make([]uint, len(bookIds))
	for i := range bookIds {
		id, _ := strconv.ParseUint(bookIds[i], 10, 32)
		ids[i] = uint(id)
	}

	return ids
}

func (b *AuthorPostgres) GetBook(bookId uint) (models.Book, error) {
	var book models.Book
	if err := b.db.First(&book, bookId).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}
