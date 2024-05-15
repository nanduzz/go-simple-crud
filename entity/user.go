package entity

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User is a struct that represents a user
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Username string
	Password string
}

// NewUser is a function that creates a new user
func NewUser(id primitive.ObjectID, username, password string) *User {
	if id.IsZero() {
		id = primitive.NewObjectID()
	}
	return &User{
		ID:       id,
		Username: username,
		Password: password,
	}
}

// String is a function that returns a string representation of a user
func (u *User) String() string {
	return fmt.Sprintf("ID: %s\nUsername: %s", u.ID, u.Username)
}
