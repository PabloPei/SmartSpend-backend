package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/PabloPei/SmartSpend-backend/conf"
	"github.com/PabloPei/SmartSpend-backend/internal/handlers/users"
	users2 "github.com/PabloPei/SmartSpend-backend/internal/repositories/users"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(cfg conf.ApiServerConfig, db *sql.DB) *APIServer {
	return &APIServer{
		addr: fmt.Sprintf("%s:%s", cfg.PublicHost, cfg.Port),
		db:   db,
	}
}

func (s *APIServer) Run() {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	subrouter.HandleFunc("/home", users.Home).Methods("POST", "GET")

	userStore := users2.NewStore(s.db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Server running on", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, router))

}
