package repo

import (
	"context"
	"fmt"
	"follow-service/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type FollowRepo struct {
	client *mongo.Client
}

func NewRepo() (*FollowRepo, error) {
	followRepo := &FollowRepo{}

	//mongoUri := os.Getenv("MONGODB_URI")
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

	followRepo.client = client

	return followRepo, nil
}

func (followRepo *FollowRepo) CloseDB() error {
	err := followRepo.client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}
	return err
}

func (followRepo *FollowRepo) Follow(userId string, followerId string) error {
	follow := data.Follow{
		FollowerID: followerId,
		FolloweeID: userId,
	}
	collection := followRepo.client.Database("follows").Collection("follows")
	insertResult, err := collection.InsertOne(context.TODO(), follow)
	if err != nil {
		return err
	}
	fmt.Println("inserted follow: ", insertResult)
	return nil

}

func (followRepo *FollowRepo) FollowRequest(userId string, followerId string) error {
	follow := data.FollowRequest{
		FollowerID:    followerId,
		FolloweeID:    userId,
		RequestStatus: 0,
	}
	collection := followRepo.client.Database("follows").Collection("followRequests")
	insertResult, err := collection.InsertOne(context.TODO(), follow)
	if err != nil {
		return err
	}
	fmt.Println("inserted post: ", insertResult)
	return nil

}

func (followRepo *FollowRepo) GetAllFollows() []*data.Follow {
	var follows []*data.Follow
	collection := followRepo.client.Database("follows").Collection("follows")

	cur, err := collection.Find(context.TODO(), bson.D{{}})
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.Follow

		err := cur.Decode(&elem)
		fmt.Println(elem)
		if err != nil {
			log.Fatal(err)
		}
		follows = append(follows, &elem)
	}
	return follows
}

func (followRepo *FollowRepo) GetAllFollowers(id string) []*data.Follow {
	var followers []*data.Follow
	collection := followRepo.client.Database("follows").Collection("follows")
	filter := bson.D{{"followeeId", id}}

	cur, err := collection.Find(context.TODO(), filter)
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.Follow

		err := cur.Decode(&elem)
		fmt.Println(elem)
		if err != nil {
			log.Fatal(err)
		}
		followers = append(followers, &elem)
	}
	return followers
}

func (followRepo *FollowRepo) GetAllFollowed(id string) []*data.Follow {
	var followedAccs []*data.Follow
	collection := followRepo.client.Database("follows").Collection("follows")
	filter := bson.D{{"followerId", id}}

	cur, err := collection.Find(context.TODO(), filter)
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.Follow

		err := cur.Decode(&elem)
		fmt.Println(elem)
		if err != nil {
			log.Fatal(err)
		}
		followedAccs = append(followedAccs, &elem)
	}
	return followedAccs
}

func (followRepo *FollowRepo) AcceptFollower(userId string, followerId string) error {
	collection := followRepo.client.Database("follows").Collection("followRequests")
	filter := bson.D{{"$and", bson.A{bson.D{{"followerId", followerId}}, bson.D{{"followeeId", userId}}}}}
	update := bson.D{{"$set", bson.D{{"requestStatus", 2}}}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(result, "result")

	return nil

}

func (followRepo *FollowRepo) GetAllFollowRequests(id string) []*data.FollowRequest {
	var requests []*data.FollowRequest
	collection := followRepo.client.Database("follows").Collection("followRequests")
	filter := bson.D{{"$and", bson.A{bson.D{{"followeeId", id}}, bson.D{{"requestStatus", 0}}}}}

	cur, err := collection.Find(context.TODO(), filter)
	fmt.Println(cur)
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) {
		var elem data.FollowRequest

		err := cur.Decode(&elem)
		fmt.Println(elem)
		if err != nil {
			log.Fatal(err)
		}
		requests = append(requests, &elem)
	}
	return requests
}

func (followRepo *FollowRepo) IsFollowing(followerId string, followeeId string) interface{} {
	collection := followRepo.client.Database("follows").Collection("follows")

	filter := bson.D{{"$and", bson.A{bson.D{{"followerId", followerId}}, bson.D{{"followeeId", followeeId}}}}}

	var result data.Follow
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result, "single result from follows")
	if result.FollowerID == "" {
		return false
	} else {
		return true
	}

}
