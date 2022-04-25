package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	err = server.CloseDB()
	if err != nil {
		log.Fatal(err)
	}

	router.HandleFunc("/user", server.CreateUserHandler).Methods("POST")

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
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
