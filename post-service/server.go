package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
	"post-service/repo"
)

type PostServer struct {
	postRepo *repo.PostRepo
}

func (postServer *PostServer) CloseDB() error {
	return postServer.postRepo.CloseDB()
}

func (postServer *PostServer) CreatePostHandler(writer http.ResponseWriter, request *http.Request) {

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
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return

	}

	err = postServer.postRepo.CreatePost(rt.Content, rt.UserID)
	if err != nil {
		fmt.Println(err)
	}

}

func (postServer *PostServer) GetAllPostsHandler(writer http.ResponseWriter, request *http.Request) {
	posts := postServer.postRepo.GetAllPosts()

	renderJSON(writer, posts)

}

func (postServer *PostServer) GetAllPostsFromUserHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	user := vars["user"]
	if user == "" {
		http.Error(writer, "missing user id", http.StatusMethodNotAllowed)
	}
	posts := postServer.postRepo.GetPostsFromUser(user)
	renderJSON(writer, posts)
}

func (postServer *PostServer) LikeAPostHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	postId := vars["id"]
	fmt.Println(postId, "this is post id")
	if postId == "" {
		http.Error(writer, "missing user id", http.StatusMethodNotAllowed)
	}
	err := postServer.postRepo.LikeAPost(postId)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "error :)", http.StatusInternalServerError)
	}
	renderJSON(writer, "succes")
}

func (postServer *PostServer) DislikeAPostHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	postId := vars["id"]
	fmt.Println(postId, "this is post id")
	if postId == "" {
		http.Error(writer, "missing user id", http.StatusMethodNotAllowed)
	}
	err := postServer.postRepo.DislikeAPost(postId)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "error :)", http.StatusInternalServerError)
	}
	renderJSON(writer, "succes")
}

func NewPostServer() (*PostServer, error) {
	postRepo, err := repo.NewRepo()
	if err != nil {
		return nil, err
	}
	return &PostServer{
		postRepo: postRepo,
	}, nil

}
