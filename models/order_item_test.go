package models

import "testing"

func TestOrderItem_Validate(t *testing.T) {
	type fields struct {
		ID               int
		Product          Product
		RequiredQuantity float32
		ActualQuantity   float32
		BatchNumber      string
		Comment          string
		IsSubmitted      bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid item",
			fields: fields{
				Product: Product{
					ID:   0,
					Name: "Fish",
					Unit: "kg",
				},
				RequiredQuantity: 30.42,
				BatchNumber:      "BatchNumber",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oi := OrderItem{
				ID:               tt.fields.ID,
				Product:          tt.fields.Product,
				RequiredQuantity: tt.fields.RequiredQuantity,
				ActualQuantity:   tt.fields.ActualQuantity,
				BatchNumber:      tt.fields.BatchNumber,
				Comment:          tt.fields.Comment,
				IsSubmitted:      tt.fields.IsSubmitted,
			}
			if err := oi.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("OrderItem.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
