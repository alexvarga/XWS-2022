package data

type Interest struct {
	ID    string `bson:"_id,omitempty"`
	Title string `bson:"title"`
	Info  string `bson:"info"`
}
