package validation

import (
	"github.com/go-playground/validator/v10"
	"ketra-back/models"
)

func ValidateTicket(ticket models.Ticket) error {
	validate := validator.New()
	return validate.Struct(ticket)
}
