package validation

import (
	"github.com/go-playground/validator/v10"
	"ketra-back/models"
)

// Регистрация кастомного валидатора
func RegisterCustomValidators(v *validator.Validate) {
	// Регистрируем кастомную валидацию для поля ID
	v.RegisterValidation("no_id", func(fl validator.FieldLevel) bool {
		// Проверяем, что поле ID равно 0
		return fl.Field().Uint() == 0
	})
	v.RegisterValidation("no_status", func(fl validator.FieldLevel) bool {
		return !fl.Field().Bool()
	})
}

// Валидация тикета
func ValidateTicket(ticket *models.Ticket) error {
	// Создаем новый валидатор
	validate := validator.New()

	// Регистрируем кастомные валидаторы
	RegisterCustomValidators(validate)

	// Выполняем валидацию
	return validate.Struct(ticket)
}
