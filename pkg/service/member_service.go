package service

import (
	"github.com/aalmat/bookstore/models"
	"github.com/aalmat/bookstore/pkg/repository"
)

type MemberService struct {
	repo repository.Member
}

func (b *MemberService) Create(authorId uint, book models.Book) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (b *MemberService) GetBook(bookId uint) (models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *MemberService) Update(bookId uint, update models.UpdateBook) error {
	//TODO implement me
	panic("implement me")
}

func (b *MemberService) Delete(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (b *MemberService) GetAuthorBooks(id uint) ([]models.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberService(repo repository.Member) *MemberService {
	return &MemberService{repo}
}

func (b *MemberService) GetClientBooks(id uint) ([]models.Book, error) {
	return b.repo.GetClientBooks(id)
}

func (b *MemberService) OrderBook(userId, bookId uint) (uint, error) {
	return b.repo.OrderBook(userId, bookId)
}

func (b *MemberService) DeleteOrder(userId, bookId uint) error {
	return b.repo.DeleteOrder(userId, bookId)
}
