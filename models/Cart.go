package models

type Cart struct {
	ID     uint   `json:"id"`
	UserId uint   `json:"user_id" gorm:"ForeignKey:User.ID"`
	BookId string `json:"books"` //stores list of id's as string
}
