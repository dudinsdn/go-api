package main

import (
	"log"
	"my/go-api/config"
	"my/go-api/internal/handler"
	"my/go-api/internal/repository"
	"my/go-api/internal/usecase"
	"net/http"
)

func main() {
	cfg := config.Load()
	db := config.InitDB(cfg.DB)
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUC)
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.HandleGetUsers(w, r)
		case http.MethodPost:
			userHandler.HandleCreateUser(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	log.Printf("Server running at %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(cfg.Server.Port, nil))
}
