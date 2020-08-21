package server

import (
	"alex/fishorder-api-v3/app2/storage"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// APIServer ...
type APIServer struct {
	config  *Config
	router  *mux.Router
	storage *storage.Storage
}

// New ...
func New(config *Config) (*APIServer, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}
	return &APIServer{
		config:  config,
		router:  mux.NewRouter(),
		storage: config.Storage,
	}, nil
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.storage.Connect(); err != nil {
		return fmt.Errorf("Failed to connect a storage: %v", err)
	}
	s.confiureRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) confiureRouter() {

	s.router.PathPrefix("/x/").Handler(http.StripPrefix("/x/", http.FileServer(http.Dir("./web/"))))

	api := s.router.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/products", s.HandleGetAllProducts()).Methods("GET")
	api.HandleFunc("/products", s.HandleAddProduct()).Methods("POST")
	api.HandleFunc("/orders", s.HandleAddOrder()).Methods("POST")
	api.HandleFunc("/orders/{date}", s.HandleGetOrdersByDate()).Methods("GET")
	api.HandleFunc("/order-items", s.HandleUpdateItem()).Methods("POST")
	api.HandleFunc("/update-order", s.HandleUpdateOrder()).Methods("POST")

	/*
		// api.HandleFunc("/order/{id}", s.HandleGetOrder()).Methods("GET")
	*/

}
