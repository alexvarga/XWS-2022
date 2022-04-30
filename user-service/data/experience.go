package data

type Experience struct {
	ID          string `bson:"_id,omitempty"`
	DateStarted string `bson:"dateStarted"`
	DateEnded   string `bson:"dateEnded"`
	Title       string `bson:"title"`
	Info        string `bson:"info"`
}
