package cloudplatform

import (
	snq "github.com/Luke-Vear/nettaton/pkg/subnetquiz"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// User struct contains all data about a user.
type User struct {
	ID     string            `json:"id"`
	Status string            `json:"status"`
	Marks  map[string]*Marks `json:"marks"`

	// HashedPassword is generated for new users or
	// read from the database for existing ones.
	HashedPassword string `json:"passwordHash"`

	// ClearTextPassword is submitted by the client.
	ClearTextPassword string `json:"-"`
}

// Marks tracks correct answers and overall attempts for a question kind.
type Marks struct {
	Attempts int
	Correct  int
}

// NewUser returns a *User with all question types initialised.
func NewUser(id string) *User {
	marks := make(map[string]*Marks)
	for k := range snq.Questions {
		marks[k] = &Marks{}
	}
	return &User{
		ID:    id,
		Marks: marks,
	}
}

// Create a user in the database (if none exists). Create will attempt to read
// the user from the database (this will cause no error if user is not found)
// it will then check the Status field of the User object, if the status field
// is empty after the read, the user doesn't exist in the database, so create
// with the submitted password hashed.
func (u *User) Create() error {
	if err := u.Read(); err != nil {
		return err
	}
	if u.Status != "" {
		return ErrUserAlreadyExists
	}
	u.Status = "new"
	pwh, err := genPwHash(u.ClearTextPassword)
	if err != nil {
		return err
	}
	u.HashedPassword = pwh
	return u.Update()
}

// Read user from database, if the user.ID is not found in the database,
// the User's fields won't be updated.
func (u *User) Read() error {
	query := &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(u.ID)},
		},
	}
	result, err := db.GetItem(query)
	if err != nil {
		return err
	}
	if err := dynamodbattribute.UnmarshalMap(result.Item, u); err != nil {
		return err
	}
	return nil
}

// Update put's the User fields to the database.
func (u *User) Update() error {
	avm, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return err
	}
	query := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      avm,
	}
	if _, err = db.PutItem(query); err != nil {
		return err
	}
	return nil
}

// Delete removes the user from the database.
func (u *User) Delete() error {
	query := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(u.ID)},
		},
	}
	if _, err := db.DeleteItem(query); err != nil {
		return err
	}
	return nil
}

// Login validates the users submitted ClearTextPassword against the
// HashedPassword from the database, then creates a JWT with some standard
// claims, returns the signed token as a string.
func (u *User) Login() (string, error) {
	return login(u)
}
