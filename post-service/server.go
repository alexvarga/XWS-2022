package main

import (
	"fmt"
	"mime"
	"net/http"
	"xws/post-service/repo"
)

type PostServer struct {
	postRepo *repo.PostRepo
}

func (postServer *PostServer) CloseDB() error {
	return postServer.postRepo.CloseDB()
}

func (postServer *PostServer) CreatePost(writer http.ResponseWriter, request *http.Request) {

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

func NewPostServer() (*PostServer, error) {
	postRepo, err := repo.NewRepo()
	if err != nil {
		return nil, err
	}
	return &PostServer{
		postRepo: postRepo,
	}, nil

}
