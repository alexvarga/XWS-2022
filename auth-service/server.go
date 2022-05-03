package main

import (
	"net/http"
	"xws/user-service/repo"
)

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

func (userServer *UserServer) LoginHandler(writer http.ResponseWriter, request *http.Request) {

}

func (userServer *UserServer) RegisterHandler(writer http.ResponseWriter, request *http.Request) {

}

func (userServer *UserServer) LogoutHandler(writer http.ResponseWriter, request *http.Request) {

}
