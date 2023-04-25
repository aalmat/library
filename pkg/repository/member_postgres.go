package repository

import (
	"errors"
	"fmt"
	"github.com/aalmat/bookstore/models"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type MemberPostgres struct {
	db *gorm.DB
}

func NewMemberPostgres(db *gorm.DB) *MemberPostgres {
	return &MemberPostgres{db}
}

func (b *MemberPostgres) OrderBook(userId, bookId uint) (uint, error) {
	cart, err := b.GetCart(userId)

	if err == gorm.ErrRecordNotFound {
		logrus.Println("asdfasdf")
		cart, err = b.CreateCart(userId)
		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	if cart.BookId != "" {
		cart.BookId += ","
	}
	cart.BookId += strconv.FormatUint(uint64(bookId), 10)

	err = b.db.Save(&cart).Error
	if err != nil {
		return 0, err
	}

	return cart.ID, nil
}

func (b *MemberPostgres) GetCart(userId uint) (models.Cart, error) {
	var cart models.Cart
	err := b.db.Where("user_id = ?", userId).First(&cart).Error

	if err == gorm.ErrRecordNotFound {
		return models.Cart{}, err
	}
	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (b *MemberPostgres) CreateCart(userId uint) (models.Cart, error) {
	var cart models.Cart
	cart.UserId = userId
	cart.BookId = ""
	if err := b.db.Select("user_id", "book_id").Create(&cart).Error; err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func (b *MemberPostgres) DeleteOrder(userId, bookId uint) error {
	cart, err := b.GetCart(userId)
	if err != nil {
		return err
	}

	if cart.BookId == "" {
		return errors.New("you didn't ordered book yet")
	}

	str := strconv.FormatUint(uint64(bookId), 10)

	res := strings.Replace(cart.BookId, str, "", 1)
	if res == cart.BookId {
		return errors.New(fmt.Sprintf("you haven't ordered book with id %d", bookId))
	}
	if res != "" {

		res = strings.ReplaceAll(res, ",,", ",")
		if string(res[0]) == "," {
			res = res[1:]
		}
	}
	cart.BookId = res
	if err := b.db.Save(&cart).Error; err != nil {
		return err
	}
	return nil

}

func (b *MemberPostgres) GetClientBooks(id uint) ([]models.Book, error) {
	//var bookIds []uint

	var cart models.Cart
	if err := b.db.Where("user_id = ?", id).First(&cart).Error; err != nil {
		return nil, err
	}
	bookIds := strings.Split(cart.BookId, ",")
	ids := getBookId(bookIds)
	var books []models.Book
	if err := b.db.Find(&books, ids).Error; err != nil {
		return nil, err
	}

	return books, nil
}
