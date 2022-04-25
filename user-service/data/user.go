package data

type User struct {
	ID        int    `bson:"id"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
}
type ResponseId struct {
	Id int `bson:"id"`
}
