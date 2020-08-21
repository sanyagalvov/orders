package storage

import validation "github.com/go-ozzo/ozzo-validation"

// Config contains all information needed for connecting to DB.
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

// Validate ...
func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Host, validation.Required),
		validation.Field(&c.Port, validation.Required),
		validation.Field(&c.User, validation.Required),
		validation.Field(&c.Password, validation.Required),
		validation.Field(&c.DBname, validation.Required),
	)
}
