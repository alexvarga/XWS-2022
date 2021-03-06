package main

import (
	"context"
	"github.com/gorilla/handlers"
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
	server, err := NewPostServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	defer server.CloseDB()

	headers := handlers.AllowedHeaders([]string{"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Origin", "Accept", "Authorization", "Access-Control-Allow-Origin"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	credentials := handlers.AllowCredentials()

	router.HandleFunc("/post", server.CreatePostHandler).Methods("POST")
	router.HandleFunc("/post/{id}", server.GetPostByIdHandler).Methods("GET")
	router.HandleFunc("/posts", server.GetAllPostsHandler).Methods("GET")
	router.HandleFunc("/posts/{user}", server.GetAllPostsFromUserHandler).Methods("GET")
	router.HandleFunc("/post/like/{id}", server.LikeAPostHandler).Methods("PUT")
	router.HandleFunc("/post/dis/{id}", server.DislikeAPostHandler).Methods("PUT")
	router.HandleFunc("/post/comment/{postId}", server.LeaveACommentHandler).Methods("POST")

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
