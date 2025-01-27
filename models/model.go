package models

import (
	"fmt"
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

func DeleteTicketByID(id uint) error {
	result := db.DB.Delete(&Ticket{}, id)
	if result.Error != nil {
		log.Printf("Ошибка при удалении заявки с ID %d: %v", id, result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("Заявка с ID %d не найдена", id)
		return fmt.Errorf("заявка с ID %d не найдена", id)
	}

	return nil
}

func GetTotalTicketCount() (int64, error) {
	var count int64
	if err := db.DB.Model(&Ticket{}).Count(&count).Error; err != nil {
		log.Printf("Ошибка при получении общего количества заявок: %v", err)
		return 0, err
	}
	return count, nil
}

func GetTickets(limit, offset int) ([]Ticket, error) {
	var tickets []Ticket
	if err := db.DB.Limit(limit).Offset(offset).Order("id DESC").Find(&tickets).Error; err != nil {
		log.Printf("Ошибка при получении всех заявок: %v", err)
		return nil, err
	}
	return tickets, nil
}
