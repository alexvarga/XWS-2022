package data

type FollowRequest struct {
	ID            string `bson:"_id,omitempty"`
	FollowerID    string `bson:"followerId"`
	FolloweeID    string `bson:"followeeId"`
	RequestStatus int    `bson:"requestStatus"`
}
