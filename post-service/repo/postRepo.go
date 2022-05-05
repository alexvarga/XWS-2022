package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"xws/post-service/data"
)

type PostRepo struct {
	client *mongo.Client
}

func NewRepo() (*PostRepo, error) {
	postRepo := &PostRepo{}

	//mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + "localhost:27017" + "/?connect=direct")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	postRepo.client = client

	return postRepo, nil
}

func (postRepo *PostRepo) CloseDB() error {
	err := postRepo.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
	return err
}

func (postRepo *PostRepo) CreatePost(content []byte, id string) error {
	post := data.Post{
		Content:   content,
		UserID:    id,
		Published: time.Now(),
	}

	collection := postRepo.client.Database("posts").Collection("posts")

	insertResult, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	fmt.Println("inserted post: ", insertResult)
	return nil
}
