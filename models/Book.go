package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title" binding:"required"`
	Isbn   string `json:"isbn" binding:"required"`
	Genres Genre  `json:"genres" binding:"required"`
	UserId uint   `json:"user_id" gorm:"ForeignKey:User.ID"`
}

type UpdateBook struct {
	Title string `json:"title"`
	Isbn  string `json:"isbn"`
}

type Genre int

const (
	Documental = iota + 1
	Historical
	Action
)

func (u *UpdateBook) Validate() error {
	if u.Title == "" && u.Isbn == "" {
		return errors.New("Nothing to change")
	}
	return nil
}
