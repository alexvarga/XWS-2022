package data

type User struct {
	ID          string  `bson:"_id,omitempty"`
	Email       string  `bson:"email"`
	Password    string  `bson:"password"`
	FirstName   string  `bson:"firstName"`
	LastName    string  `bson:"lastName"`
	Age         string  `bson:"age"`
	Gender      *Gender `bson:"gender"`
	PhoneNumber string  `bson:"phoneNumber"`
	Bio         string  `bson:"bio"`

	PrivateProfile bool    `bson:"privateProfile"`
	Following      []*User `bson:"following"`
}
type ResponseId struct {
	Id string `bson:"_id"`
}
