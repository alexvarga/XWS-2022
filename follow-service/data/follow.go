package data

type Follow struct {
	ID         string `bson:"_id,omitempty"`
	FollowerID string `bson:"followerId"`
	FolloweeID string `bson:"followeeId"`
}
