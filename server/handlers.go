package server

import (
	"alex/fishorder-api-v3/app2/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// HandleGetAllProducts ...
func (s *APIServer) HandleGetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		p, err := s.storage.SelectAllProducts()
		if err != nil {
			logrus.Fatal(err)
		}
		json.NewEncoder(w).Encode(p)
	}
}

// HandleAddProduct ...
func (s *APIServer) HandleAddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var p models.Product
		_ = json.NewDecoder(r.Body).Decode(&p)
		if err := p.Validate(); err != nil {
			json.NewEncoder(w).Encode(&models.Product{})
			return
		}
		err := s.storage.InsertProduct(&p)
		if err != nil {
			logrus.Fatal(err)
		}
		json.NewEncoder(w).Encode(p)
	}
}

// HandleUpdateItem ...
func (s *APIServer) HandleUpdateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var p models.OrderItem
		_ = json.NewDecoder(r.Body).Decode(&p)
		if err := p.Validate(); err != nil {
			json.NewEncoder(w).Encode(&models.Product{})
			return
		}
		err := s.storage.UpdateOrderItem(&p)
		if err != nil {
			logrus.Fatal(err)
		}
		json.NewEncoder(w).Encode(p)
	}
}

// HandleAddOrder ...
func (s *APIServer) HandleAddOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var p models.Order
		_ = json.NewDecoder(r.Body).Decode(&p)

		fmt.Println(p)

		if err := p.Validate(); err != nil {
			logrus.Error(err)
			json.NewEncoder(w).Encode(&models.Order{})
			return
		}
		err := s.storage.InsertOrder(&p)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(p)
	}
}

// HandleGetOrdersByDate ...
func (s *APIServer) HandleGetOrdersByDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		dateString := params["date"]
		date, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			log.Fatal(err)
		}

		p, err := s.storage.SelectOrdersByDate(date)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(p)
	}
}

// HandleUpdateOrder ...
func (s *APIServer) HandleUpdateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var p models.Order
		_ = json.NewDecoder(r.Body).Decode(&p)
		if err := p.Validate(); err != nil {
			json.NewEncoder(w).Encode(&models.Product{})
			return
		}
		err := s.storage.UpdateOrder(&p)
		if err != nil {
			logrus.Fatal(err)
		}
		json.NewEncoder(w).Encode(p)
	}
}
