package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	router := mux.NewRouter()
	router.StrictSlash(true)
	server, err := NewFollowServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer server.CloseDB()

	router.HandleFunc("/follow/{user}/{follower}", server.FollowHandler).Methods("POST")
	router.HandleFunc("/follow/request/{user}/{follower}", server.FollowRequestHandler).Methods("POST")
	router.HandleFunc("/follows", server.GetAllFollowsHandler).Methods("GET")
	router.HandleFunc("/followers/{userId}", server.GetAllFollowersHandler).Methods("GET") //svi koji nas prate
	router.HandleFunc("/following/{userId}", server.GetAllFollowedHandler).Methods("GET")  //svi koje pratimo
	router.HandleFunc("/follow/{user}/{follower}", server.AcceptFollowerHandler).Methods("PUT")
	router.HandleFunc("/follower/requests/{user}", server.GetAllFollowRequestsHandler).Methods("GET")

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
