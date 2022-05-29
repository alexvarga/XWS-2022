package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mime"
	"net/http"
	"user-service/data"

	"user-service/repo"
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

func (userServer *UserServer) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(r.Body)
	if err != nil {
		//tracer.LogError(span, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := userServer.userRepo.CreateUser(rt.Email, rt.Password, rt.FirstName, rt.LastName, rt.Age, rt.Gender, rt.Bio, rt.PhoneNumber)
	if id != "" {
		fmt.Println(id)
		renderJSON(w, data.ResponseId{Id: id})
	} else if id == "" {
		http.Error(w, "user with this email already exists", http.StatusMethodNotAllowed)

	}

}

func (userServer *UserServer) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := userServer.userRepo.GetAllUsers()
	renderJSON(w, users)
}

func (userServer *UserServer) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "missing id", http.StatusMethodNotAllowed)
	}
	user := userServer.userRepo.GetUserById(id)
	fmt.Println(user, "ovo je user")
	renderJSON(w, user)
}

func (userServer *UserServer) UpdateUserHandler(writer http.ResponseWriter, request *http.Request) {
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

	ru, err := decodeBody(request.Body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err1 := userServer.userRepo.EditUser(ru.Email, ru.FirstName, ru.LastName, ru.Age, ru.Gender, ru.PhoneNumber, ru.Bio, ru.PrivateProfile)
	if err1 != nil {
		http.Error(writer, err.Error(), http.StatusMethodNotAllowed)
	}
	renderJSON(writer, "success")
}

func (userServer *UserServer) SearchUsersHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	search := vars["search"]

	users := userServer.userRepo.SearchUsers(search)
	renderJSON(writer, users)

}

func (userServer *UserServer) AddExperienceHandler(writer http.ResponseWriter, request *http.Request) {
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
		//tracer.LogError(span, err)
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = userServer.userRepo.AddExperience(rt.Email, rt.Experience)
	if err != nil {
		log.Fatal(err)
	}

}

func (userServer *UserServer) UpdateExperienceHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	if id == "" {
		http.Error(writer, "missing id", http.StatusMethodNotAllowed)
	}

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

	err = userServer.userRepo.UpdateExperience(rt.Email, id, rt.Experience)
	if err != nil {
		log.Fatal(err)
	}

}

func (userServer *UserServer) GetUserIdByEmailHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]
	if email == "" {
		http.Error(writer, "missing id", http.StatusMethodNotAllowed)
	}
	user := userServer.userRepo.GetUserByEmail(email)
	fmt.Println(user, "ovo je user")
	renderJSON(writer, user.ID)
}

func (userServer *UserServer) GetUserByEmailHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	email := vars["email"]
	if email == "" {
		http.Error(writer, "missing id", http.StatusMethodNotAllowed)
	}
	user := userServer.userRepo.GetUserByEmail(email)
	fmt.Println(user, "ovo je user")
	renderJSON(writer, user)
}

//func (s *UserServer) CreateUser(ctx context.Context, in *usr.CreateUserRequest) (*usr.User, error) {
//
//}
