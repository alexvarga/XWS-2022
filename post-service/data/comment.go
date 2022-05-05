package data

type Comment struct {
	ID       string `bson:"_id,omitempty"`
	PostID   string `bson:"postId"`
	UserID   string `bson:"userId"`
	Text     string `bson:"text"`
	Likes    int    `bson:"likes"`
	Dislikes int    `bson:"dislikes"`
}
