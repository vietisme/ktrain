package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"ktrain/cmd/api/user-api/handler"
	middleware2 "ktrain/cmd/api/user-api/middleware"
	"ktrain/cmd/repository"
	"ktrain/pkg/storage"
	"log"
	"net/http"
)

func main() {
	psqlDB, err := storage.NewPSQLManager()
	if err != nil {
		log.Fatalf("Error when connecting database, err: %v", err)
		return
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("Content-Type", "application/json"))

		userRepository := repository.NewUserRepository(psqlDB)

		//Authenticate
		r.Use(middleware2.NewDBTokenAuth(userRepository).Handle())

		//API handlers
		userHandler := handler.NewUserHandler(userRepository)
		r.Get("/me", userHandler.GetMyProfile)
	})

	http.ListenAndServe(":8080", r)
}
