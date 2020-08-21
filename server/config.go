package server

import (
	"alex/fishorder-api-v3/app2/storage"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Config ...
type Config struct {
	BindAddr string
	Storage  *storage.Storage
}

// Validate ...
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.BindAddr, validation.Required),
		validation.Field(&c.Storage, validation.Required),
	)
}
