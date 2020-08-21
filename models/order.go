package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Order ...
type Order struct {
	ID           int         `json:"id,omitempty"`
	Recipient    string      `json:"recipient"`
	ShippingDate time.Time   `json:"date"`
	Items        []OrderItem `json:"items"`
	Comment      string      `json:"comment,omitempty"`
	IsSubmitted  bool        `json:"is_submitted"`
}

// Validate ...
func (o Order) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Recipient, validation.Required),
		validation.Field(&o.ShippingDate, validation.Required),
		validation.Field(&o.Items, validation.Required),
		validation.Field(&o.IsSubmitted),
	)
}
