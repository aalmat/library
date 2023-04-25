package repository

import (
	"errors"
	"github.com/aalmat/bookstore/models"
	"github.com/jinzhu/gorm"
	"time"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (a *AuthPostgres) CreateUser(user models.User) (uint, error) {

	var err error
	user.Password, err = hashingPassword(user.Password)
	if err != nil {
		return 0, err
	}

	if err := a.db.First(&user, "email=$1", user.Email).Error; err == nil {
		return 0, errors.New("email already registered")
	}

	user.CreatedAt, user.UpdatedAt = time.Now(), time.Now()

	a.db.Select("name", "email", "password", "created_at", "updated_at", "user_role").Create(&user)

	return user.ID, nil
}

func (a *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User

	//fmt.Println("Email: ", email)
	//fmt.Println("Password: ", password)

	if err := a.db.First(&user, "email=$1", email).Error; err != nil {
		return models.User{}, err
	}
	if err := compare(password, user.Password); err != nil {
		return models.User{}, err
	}

	return user, nil
}
