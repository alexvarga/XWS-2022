package data

import "time"

type Post struct {
	ID        string    `bson:"_id,omitempty"`
	Content   []byte    `bson:"content"` //for now
	UserID    string    `bson:"userId"`
	Published time.Time `bson:"published"`
	Likes     int       `bson:"likes"`
	Dislikes  int       `bson:"dislikes"`
	Comments  []Comment `bson:"comments"`
}
