package main

import (
	"mime"
	"net/http"
	"xws/user-service/data"
	"xws/user-service/repo"
)

//calls repo

type UserServer struct {
	userRepo *repo.UserRepo
}

func NewUserServer() (*UserServer, error) {
	userRepo, err := repo.NewRepo()
	if err != nil {
		return nil, err
	}
	return &UserServer{
		userRepo: userRepo,
	}, nil

}

func (userServer *UserServer) CloseDB() error {
	return userServer.userRepo.CloseDB()
}

func (userServer *UserServer) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		//tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	//ctx := tracer.ContextWithSpan(context.Background(), span)
	rt, err := decodeBody(r.Body)
	if err != nil {
		//tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := userServer.userRepo.CreateUser(rt.Email, rt.Password, rt.FirstName, rt.LastName)
	renderJSON(w, data.ResponseId{Id: id})

}
