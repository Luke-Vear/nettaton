package cloudplatform

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	// Same table for all database queries.
	table = os.Getenv("TABLE")

	// Secret used for signing JWTs.
	secret = os.Getenv("SECRET")

	// Same db session for all database queries.
	db GetDeletePutter
)

func init() {
	db = dynamodb.New(
		session.Must(
			session.NewSession(
				&aws.Config{
					Region: aws.String(
						os.Getenv("REGION"))})))
}

type GetDeletePutter interface {
	GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error)
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
