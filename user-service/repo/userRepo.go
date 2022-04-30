package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
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

func (userRepo *UserRepo) CreateUser(email string, password string, firstName string, lastName string) string {
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

		str, err := json.Marshal(insertResult.InsertedID)

		return string(str)
	}
	return ""
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

func (userRepo *UserRepo) GetUserById(id string) data.User {
	var user data.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.D{{"_id", objectId}}
	collection := userRepo.client.Database("testDatabase").Collection("users")
	err1 := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err1 != nil {
		fmt.Println(err1)
	}
	return user
}

func (userRepo *UserRepo) GetUserByEmail(email string) data.User {
	var user data.User

	filter := bson.D{{"email", email}}
	collection := userRepo.client.Database("testDatabase").Collection("users")
	err1 := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err1 != nil {
		fmt.Println(err1)
	}
	return user
}

func (userRepo *UserRepo) EditUser(email string, firstName string, lastName string, age string, gender *data.Gender, number string, bio string, profile bool) error {

	var user data.User
	filter := bson.D{{"email", email}}

	collection := userRepo.client.Database("testDatabase").Collection("users")
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}

	update := bson.D{
		{"$set", bson.D{{"email", email}}},
		{"$set", bson.D{{"firstName", firstName}}},
		{"$set", bson.D{{"lastName", lastName}}},
		{"$set", bson.D{{"age", age}}},
		{"$set", bson.D{{"gender", gender}}},
		{"$set", bson.D{{"phoneNumber", number}}},
		{"$set", bson.D{{"bio", bio}}},
		{"$set", bson.D{{"privateProfile", profile}}},
	}

	udpateUser, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(udpateUser, "udpate user")

	return nil
	//return user
}

func (userRepo *UserRepo) SearchUsers(search string) *[]data.User {

	searchTerms := strings.Split(search, " ")
	var results []data.User
	users := userRepo.GetAllUsers() //should be only users with non-private profiles
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(searchTerms); j++ {
			if (strings.Contains(users[i].FirstName, searchTerms[j])) || (strings.Contains(users[i].LastName, searchTerms[j])) {
				results = append(results, *users[i])
			}
		}
	}

	//fmt.Println(results, "these are results")

	return &results
}

func (userRepo *UserRepo) AddExperience(email string, experience []data.Experience) error {

	collection := userRepo.client.Database("testDatabase").Collection("users")

	update := bson.D{
		{"$push", bson.D{{"experience", bson.D{
			{"_id", primitive.NewObjectID()},
			{"dateStarted", experience[0].DateStarted},
			{"dateEnded", experience[0].DateEnded},
			{"title", experience[0].Title},
			{"info", experience[0].Info},
		}}}},
	}
	filter := bson.D{{"email", email}}

	udpateUser, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(udpateUser, "udpate user experience. it is going to break :)")

	return nil
}

func (userRepo *UserRepo) UpdateExperience(email string, id string, experience []data.Experience) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	collection := userRepo.client.Database("testDatabase").Collection("users")

	filter := bson.D{{"email", email}}
	var result data.User
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err
	}

	for i := 0; i < len(result.Experience); i++ {
		fmt.Println(result.Experience[i].ID, "obj", objectId)
		if result.Experience[i].ID == id {
			result.Experience[i].Title = experience[0].Title
			result.Experience[i].Info = experience[0].Info
			result.Experience[i].DateStarted = experience[0].DateStarted
			result.Experience[i].DateEnded = experience[0].DateEnded
		}
	}

	fmt.Println(filter, "ovo je filter")

	update := bson.D{{"$set", bson.D{{"experience", result.Experience}}}}

	fmt.Println(result, "ovo je result")

	udpateUser, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(*udpateUser, "udpated")

	return nil
}
