package models

import validation "github.com/go-ozzo/ozzo-validation/v4"

// OrderItem ...
type OrderItem struct {
	ID               int     `json:"id,omitempty"`
	Product          Product `json:"product"`
	RequiredQuantity float32 `json:"rq"`
	ActualQuantity   float32 `json:"aq,omitempty"`
	BatchNumber      string  `json:"bn,omitempty"`
	Comment          string  `json:"comment,omitempty"`
	IsSubmitted      bool    `json:"is_submitted"`
}

// Validate ...
func (oi OrderItem) Validate() error {
	return validation.ValidateStruct(&oi,
		validation.Field(&oi.Product, validation.Required),
		validation.Field(&oi.RequiredQuantity, validation.Required),
		validation.Field(&oi.IsSubmitted),
	)
}
