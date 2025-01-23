package models

import (
	"ketra-back/db"
	"log"
)

type Ticket struct {
	ID          uint   `json:"id" gorm:"primaryKey" validate:"no_id"`
	FIO         string `json:"fio" binding:"required" gorm:"not null"`
	PhoneNumber string `json:"phone_number" binding:"required" gorm:"not null"`
	Email       string `json:"email" binding:"required,email" gorm:"not null"`
	Wishlist    string `json:"wishlist"`
	Status      bool   `json:"status" gorm:"default:false" validate:"no_status"`
}

// Метод для создания новой заявки в базе данных
func (ticket *Ticket) Create() error {
	if err := db.DB.Create(ticket).Error; err != nil {
		log.Printf("Ошибка при создании заявки: %v", err)
		return err
	}
	return nil
}

// Метод для получения заявки по ID
func GetTicketByID(id uint) (*Ticket, error) {
	var ticket Ticket
	if err := db.DB.First(&ticket, id).Error; err != nil {
		log.Printf("Ошибка при поиске заявки с ID %d: %v", id, err)
		return nil, err
	}
	return &ticket, nil
}

// Метод для обновления статуса заявки
func (ticket *Ticket) UpdateStatus(status bool) error {
	ticket.Status = status
	if err := db.DB.Save(ticket).Error; err != nil {
		log.Printf("Ошибка при обновлении статуса заявки с ID %d: %v", ticket.ID, err)
		return err
	}
	return nil
}

// Метод для удаления заявки по ID
func DeleteTicketByID(id uint) error {
	if err := db.DB.Delete(&Ticket{}, id).Error; err != nil {
		log.Printf("Ошибка при удалении заявки с ID %d: %v", id, err)
		return err
	}
	return nil
}

// Метод для получения всех заявок
func GetAllTickets() ([]Ticket, error) {
	var tickets []Ticket
	if err := db.DB.Find(&tickets).Error; err != nil {
		log.Printf("Ошибка при получении всех заявок: %v", err)
		return nil, err
	}
	return tickets, nil
}
