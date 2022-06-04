package main

import (
	"fmt"
	"follow-service/repo"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowServer struct {
	followRepo *repo.FollowRepo
}

func (followServer *FollowServer) CloseDB() error {
	return followServer.followRepo.CloseDB()
}

func (followServer *FollowServer) FollowHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["user"]
	followerId := vars["follower"]
	if userId == "" || followerId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	err := followServer.followRepo.Follow(userId, followerId)
	if err != nil {
		fmt.Println(err)
	}
}

func (followServer *FollowServer) FollowRequestHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["user"]
	followerId := vars["follower"]
	if userId == "" || followerId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	err := followServer.followRepo.FollowRequest(userId, followerId)
	if err != nil {
		fmt.Println(err)
	}
}

func (followServer *FollowServer) GetAllFollowsHandler(writer http.ResponseWriter, request *http.Request) {
	followers := followServer.followRepo.GetAllFollows()

	renderJSON(writer, followers)
}

func (followServer *FollowServer) GetAllFollowersHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	if userId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	followers := followServer.followRepo.GetAllFollowers(userId)

	renderJSON(writer, followers)

}

func (followServer *FollowServer) GetAllFollowedHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]
	if userId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	followed := followServer.followRepo.GetAllFollowed(userId)
	renderJSON(writer, followed)

}

func (followServer *FollowServer) AcceptFollowerHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["user"]
	followerId := vars["follower"]
	if userId == "" || followerId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	err := followServer.followRepo.AcceptFollower(userId, followerId)
	if err != nil {
		fmt.Println(err)
	}
	followServer.FollowHandler(writer, request)
}

func (followServer *FollowServer) GetAllFollowRequestsHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["user"]
	if userId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	followed := followServer.followRepo.GetAllFollowRequests(userId)
	renderJSON(writer, followed)
}

func (followServer *FollowServer) IsFollowingHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	followeeId := vars["followee"]
	followerId := vars["follower"]
	if followeeId == "" || followerId == "" {
		http.Error(writer, "missing user or follower id", http.StatusMethodNotAllowed)
	}
	isFollowing := followServer.followRepo.IsFollowing(followerId, followeeId)

	renderJSON(writer, isFollowing)

}

func NewFollowServer() (*FollowServer, error) {
	followRepo, err := repo.NewRepo()
	if err != nil {
		return nil, err
	}
	return &FollowServer{
		followRepo: followRepo,
	}, nil

}
