package data

type User struct {
	ID             int     `bson:"id"`
	Email          string  `bson:"email"`
	Password       string  `bson:"password"`
	FirstName      string  `bson:"firstName"`
	LastName       string  `bson:"lastName"`
	Age            string  `bson:"age"`
	PrivateProfile bool    `bson:"privateProfile"`
	Following      []*User `bson:"following"`
}
type ResponseId struct {
	Id int `bson:"id"`
}
