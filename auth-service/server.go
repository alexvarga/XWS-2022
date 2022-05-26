package main

import (
	"auth-service/repository"
	"auth-service/token"
	"fmt"
	"mime"
	"net/http"
)

type UserServer struct {
	userRepo *repository.UserRepo
}

func NewUserServer() (*UserServer, error) {
	userRepo, err := repository.NewRepo()
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
	contentType := request.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		//tracer.LogError(span, err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(writer, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	//ctx := tracer.ContextWithSpan(context.Background(), span)
	rt, err := decodeBody(request.Body)
	if err != nil {
		//tracer.LogError(span, err)
		fmt.Println(err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = userServer.userRepo.CheckLogin(rt.Email, rt.Password)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := token.CreateJWT(rt.ID, rt.Email)

	fmt.Println(token)
	renderJSON(writer, token)

}

func (userServer *UserServer) RegisterHandler(writer http.ResponseWriter, request *http.Request) {

	contentType := request.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(writer, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(request.Body)
	if err != nil {

		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = userServer.userRepo.CreateUser(rt.Email, rt.Password)
	if err == nil {
		renderJSON(writer, "success")
	} else if err != nil {
		http.Error(writer, "user with this email already exists", http.StatusMethodNotAllowed)

	}

}

func (userServer *UserServer) LogoutHandler(writer http.ResponseWriter, request *http.Request) {

}
