package repo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"xws/user-service/data"
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

func (userRepo *UserRepo) CreateUser(email string, password string, firstName string, lastName string) int {
	//password currently palintext

	user := data.User{

		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	collection := userRepo.client.Database("testDatabase").Collection("users")

	filter := bson.D{{"email", email}}
	var result data.User
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		insertResult, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)

		return user.ID
	}
	//fmt.Println(result)

	return -1
}

func (userRepo *UserRepo) GetAllUsers() []*data.User {
	var users []*data.User
	collection := userRepo.client.Database("testDatabase").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.User

		err := cur.Decode(&elem)
		fmt.Println(elem)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &elem)
	}

	return users
}
