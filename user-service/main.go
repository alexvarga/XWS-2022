package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//call server that calls repo

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)
	server, err := NewUserServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer server.CloseDB()

	headers := handlers.AllowedHeaders([]string{"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Origin", "Accept", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	router.HandleFunc("/user", server.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users", server.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/user/id/{email}", server.GetUserIdByEmailHandler).Methods("GET")
	router.HandleFunc("/user/email/{email}", server.GetUserByEmailHandler).Methods("GET")
	router.HandleFunc("/user/{id}", server.GetUserByIdHandler).Methods("GET")
	router.HandleFunc("/user", server.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/user/search/{search}", server.SearchUsersHandler).Methods("GET")
	router.HandleFunc("/user/experience", server.AddExperienceHandler).Methods("POST")
	router.HandleFunc("/user/experience/{id}", server.UpdateExperienceHandler).Methods("PUT")
	//router.HandleFunc("/", server.AddInterestHandler)
	//router.HandleFunc("/", server.UpdateInterestHandler)

	corsHandler := handlers.CORS(headers, methods, origins, credentials)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      corsHandler(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {

		}
	}()

	<-quit

	log.Println("Service shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped.")

}
