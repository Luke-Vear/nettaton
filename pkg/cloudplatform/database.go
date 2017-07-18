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
	Email  string            `json:"email"`
	Status string            `json:"status"`
	Marks  map[string]*Marks `json:"marks"`

	// HashedPassword is generated or read from the database.
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

// Create does
// GetUser deserializes the user data into the *User struct.
// Build query from environment and User passed in to function.
// Primary key is id passed in from User.
// Unmarshal result into User struct passed in.
// If the password is an empty string, user isn't in database.
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

// Create does
// GetUser deserializes the user data into the *User struct.
// Build query from environment and User passed in to function.
// Primary key is id passed in from User.
// Unmarshal result into User struct passed in.
// If the password is an empty string, user isn't in database.
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

// Update does
// Marshal User into attribute value map for db query.
// Build query to insert data.
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

// Delete NYI
// TODO
func (u *User) Delete() error {
	return nil
}

// Login does
// GetUser deserializes the user data into the *User struct.
// Build query from environment and User passed in to function.
// Primary key is id passed in from User.
// Unmarshal result into User struct passed in.
// If the password is an empty string, user isn't in database.
func (u *User) Login() (string, error) {
	return login(u)
}
