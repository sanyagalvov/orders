package models

import validation "github.com/go-ozzo/ozzo-validation"

// Product ...
type Product struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Unit string `json:"unit"`
}

// Validate ...
func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&p.Unit, validation.Required, validation.Length(1, 100)),
	)
}
