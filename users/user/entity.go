package user

import (
	"time"
)

// Ideally the JSON and BSON tags should be defined in an specific entity
// for the server and storage components, respectively, but this being a small
// service they are defined here for simplicity.
type User struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"password" bson:"-"`
	Hash      string    `json:"-" bson:"hash"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Credentials struct {
	Email    string
	Password string
}

type CreateService interface {
	Create(User) error
}

type AuthenticateService interface {
	Authenticate(cred Credentials) (string, error)
}

type AuthorizeService interface {
	Authorize(string) (string, error)
}

type GenerateService interface {
	GenerateJWT(string) (string, error)
}

type Repository interface {
	InsertOne(User) (interface{}, error)
	FindOne(string) (User, error)
}
