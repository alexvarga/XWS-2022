package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"post-service/data"
	"time"
)

type PostRepo struct {
	client *mongo.Client
}

func NewRepo() (*PostRepo, error) {
	postRepo := &PostRepo{}

	mongoUri := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI("mongodb://" + mongoUri + "/?connect=direct")

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

func (postRepo *PostRepo) GetAllPosts() []*data.Post {
	var posts []*data.Post
	collection := postRepo.client.Database("posts").Collection("posts")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.Post

		err := cur.Decode(&elem)
		fmt.Println(elem)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &elem)
	}
	return posts

}

func (postRepo *PostRepo) GetPostsFromUser(user string) []*data.Post {
	var posts []*data.Post

	filter := bson.D{{"userId", user}}
	collection := postRepo.client.Database("posts").Collection("posts")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.Post

		err := cur.Decode(&elem)
		fmt.Println(elem, "this is a post")
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &elem)
	}
	return posts

}

func (postRepo *PostRepo) LikeAPost(id string) error {
	collection := postRepo.client.Database("posts").Collection("posts")
	objectID, err := primitive.ObjectIDFromHex(id)
	fmt.Println("object ID: ", objectID)
	if err != nil {

	}
	filter := bson.D{{"_id", objectID}}
	update := bson.D{
		{"$inc", bson.D{
			{"likes", 1},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("increased...", result)

	return nil
}

func (postRepo *PostRepo) DislikeAPost(id string) error {
	collection := postRepo.client.Database("posts").Collection("posts")
	objectID, err := primitive.ObjectIDFromHex(id)
	fmt.Println("object ID: ", objectID)
	if err != nil {

	}
	filter := bson.D{{"_id", objectID}}
	update := bson.D{
		{"$inc", bson.D{
			{"dislikes", 1},
		}},
	}
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("decreased...", result)

	return nil
}
