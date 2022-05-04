package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"xws/auth-service/data"
)

type UserRepo struct {
	client *mongo.Client
}

func NewRepo() (*UserRepo, error) {
	userRepo := &UserRepo{}

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

	userRepo.client = client

	return userRepo, nil
}

func (userRepo *UserRepo) CloseDB() error {
	err := userRepo.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
	return err
}

func (userRepo *UserRepo) CheckLogin(email string, password string) error {

	collection := userRepo.client.Database("credentials").Collection("users")

	filter := bson.D{{"email", email}}
	fmt.Println("ovo je filter: ", filter, "#")
	var result data.User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	fmt.Println("ovo je single result", err)

	fmt.Println("ovo je result: ", result, "#")
	if err != nil {
		fmt.Println("not found user")
		return errors.New("user not found")
	} else if result.Password != password {
		fmt.Println("passwords not matching")
		return errors.New("wrong password")
	} else {
		return nil
	}

}

func (userRepo *UserRepo) CreateUser(email string, password string) error {

	user := data.User{
		Email:    email,
		Password: password,
	}
	collection := userRepo.client.Database("credentials").Collection("users")

	filter := bson.D{{"email", email}}
	var result data.User
	singleResult := collection.FindOne(context.TODO(), filter).Decode(&result)
	if singleResult != nil {
		insertResult, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)

		str, err := json.Marshal(insertResult.InsertedID)
		if err != nil {
			return err
		}

		fmt.Println(str)
		return nil
	}

	return nil
}
