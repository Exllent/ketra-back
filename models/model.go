package models

type Ticket struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FIO         string `json:"fio" gorm:"not null"`
	PhoneNumber string `json:"phone_number" gorm:"not null"`
	Email       string `json:"email" gorm:"not null"`
	Wishlist    string `json:"wishlist"`
	Status      bool   `json:"status" gorm:"default:false"`
}
