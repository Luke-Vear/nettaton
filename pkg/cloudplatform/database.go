package cloudplatform

import (
	"os"

	snq "github.com/Luke-Vear/nettaton/pkg/subnetquiz"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	// Same db session for all database queries.
	db = dynamodb.New(session.Must(session.NewSession(&aws.Config{Region: aws.String(os.Getenv("REGION"))})))

	// Same table for all database queries.
	table = os.Getenv("TABLE")
)

// User struct contains all data about a user.
type User struct {
	UserID   string                    `json:"userID"`
	Password string                    `json:"password"`
	Email    string                    `json:"email"`
	Status   string                    `json:"status"`
	Scores   map[string]*QuestionScore `json:"scores"`
}

// QuestionScore tracks correct answers and overall attempts for a question kind.
type QuestionScore struct {
	Attempts int
	Correct  int
}

// NewUser returns a *User with all question types initialised.
func NewUser(userID string) *User {
	scores := make(map[string]*QuestionScore)
	for k := range snq.QuestionFuncMap {
		scores[k] = &QuestionScore{}
	}
	return &User{
		UserID: userID,
		Scores: scores,
	}
}

// GetUser deserializes the user data into the *User struct.
func GetUser(u *User) error {

	// Build query from environment and User passed in to function.
	query := &dynamodb.GetItemInput{

		TableName: aws.String(table),

		// Primary key is userID passed in from User.
		Key: map[string]*dynamodb.AttributeValue{
			"userID": {S: aws.String(u.UserID)},
		},
	}

	result, err := db.GetItem(query)
	if err != nil {
		return err
	}

	// Unmarshal result into User struct passed in.
	if err := dynamodbattribute.UnmarshalMap(result.Item, u); err != nil {
		return err
	}

	// If the password is an empty string, user isn't in database.
	if u.Password == "" {
		return ErrUserNotFoundInDatabase
	}

	return nil
}

// PutUser puts serializes the *User into the database.
func PutUser(u *User) error {

	// If new user, replace password with password hash.
	if u.Status == "" {
		err := GenPasswordHash(u, u.Password)
		if err != nil {
			return err
		}
		u.Status = "new"
	}

	// Marshal User into attribute value map for db query.
	avm, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return err
	}

	// Build query to insert data.
	query := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      avm,
	}

	if _, err = db.PutItem(query); err != nil {
		return err
	}

	return nil
}
