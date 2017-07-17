package platform

import (
	"os"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/do"
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

// GetUser deserializes the user data into the *User struct.
func GetUser(u *do.User) error {

	// // If UserID field is empty, we don't have PK required for query.
	// if u.UserID == "" {
	// 	return ErrUserNotSpecified
	// }

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
func PutUser(u *do.User) error {

	// // If UserID field is empty, we don't have PK required for query.
	// if u.UserID == "" {
	// 	return ErrUserNotSpecified
	// }

	// If new user, replace password with password hash.
	if u.Status == "" {
		err := auth.GenPasswordHash(u, u.Password)
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
