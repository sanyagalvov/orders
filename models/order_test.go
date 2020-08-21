package models

import (
	"testing"
	"time"
)

func TestOrder_Validate(t *testing.T) {
	type fields struct {
		ID           int
		Recipient    string
		ShippingDate time.Time
		Items        []OrderItem
		Comment      string
		IsSubmitted  bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "full valid order",
			fields: fields{
				ID:           0,
				Recipient:    "Maxima",
				ShippingDate: time.Date(2004, time.January, 20, 0, 0, 0, 0, time.UTC),
				Items: []OrderItem{
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid order without ID",
			fields: fields{
				Recipient:    "Maxima",
				ShippingDate: time.Date(2004, time.January, 20, 0, 0, 0, 0, time.UTC),
				Items: []OrderItem{
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "without Recipient",
			fields: fields{
				ID:           0,
				ShippingDate: time.Date(2004, time.January, 20, 0, 0, 0, 0, time.UTC),
				Items: []OrderItem{
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "without Date",
			fields: fields{
				ID:        0,
				Recipient: "Maxima",
				Items: []OrderItem{
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "without items",
			fields: fields{
				ID:           0,
				Recipient:    "Maxima",
				ShippingDate: time.Date(2004, time.January, 20, 0, 0, 0, 0, time.UTC),
			},
			wantErr: true,
		},
		{
			name: "invalid item",
			fields: fields{
				Recipient:    "Maxima",
				ShippingDate: time.Date(2004, time.January, 20, 0, 0, 0, 0, time.UTC),
				Items: []OrderItem{
					{
						Product: Product{
							ID:   0,
							Name: "Fish",
							Unit: "kg",
						},
						RequiredQuantity: 30.42,
						BatchNumber:      "BatchNumber",
					},
					{
						Product: Product{
							ID:   0,
							Name: "",
							Unit: "kg",
						},
						BatchNumber: "BatchNumber",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Order{
				ID:           tt.fields.ID,
				Recipient:    tt.fields.Recipient,
				ShippingDate: tt.fields.ShippingDate,
				Items:        tt.fields.Items,
				Comment:      tt.fields.Comment,
				IsSubmitted:  tt.fields.IsSubmitted,
			}
			if err := o.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Order.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
