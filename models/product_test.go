package models

import (
	"strings"
	"testing"
)

func TestProduct_Validate(t *testing.T) {
	type fields struct {
		ID   int
		Name string
		Unit string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Valid Product",
			fields: fields{
				ID:   0,
				Name: "Fish",
				Unit: "kg",
			},
			wantErr: false,
		},
		{
			name: "without Name",
			fields: fields{
				ID:   0,
				Unit: "kg",
			},
			wantErr: true,
		},
		{
			name: "without Unit",
			fields: fields{
				ID:   0,
				Name: "Fish",
			},
			wantErr: true,
		},
		{
			name: "short Unit and Name",
			fields: fields{
				ID:   0,
				Name: "",
				Unit: "",
			},
			wantErr: true,
		},
		{
			name: "too long Unit and Name",
			fields: fields{
				ID:   0,
				Name: strings.Repeat("#", 101),
				Unit: strings.Repeat("#", 101),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Product{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
				Unit: tt.fields.Unit,
			}
			if err := p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Product.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
