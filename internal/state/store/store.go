package store

import (
	"fmt"

	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	ErrQuestionNotFound = fmt.Errorf("questionID not found")
)

// DB represents a physical datastore.
type DB struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

// NewDB returns a instantiated DB.
func NewDB(table string, dynamo *dynamodb.DynamoDB) *DB {
	return &DB{
		table:  table,
		dynamo: dynamo,
	}
}

// GetQuestion retrieves a question in the database by primary key.
func (db *DB) GetQuestion(questionID string) (*quiz.Question, error) {

	query := &dynamodb.GetItemInput{
		TableName: aws.String(db.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(questionID)},
		},
	}

	result, err := db.dynamo.GetItem(query)
	if err != nil {
		return nil, err
	}

	var qq *quiz.Question
	if err := dynamodbattribute.UnmarshalMap(result.Item, qq); err != nil {
		return nil, err
	}

	if len(qq.IP) == 0 {
		return nil, ErrQuestionNotFound
	}

	return qq, nil
}

// UpdateQuestion creates or overwrites a question in the database by primary key.
func (db *DB) UpdateQuestion(qq *quiz.Question) error {

	avm, err := dynamodbattribute.MarshalMap(qq)
	if err != nil {
		return err
	}

	pii := &dynamodb.PutItemInput{
		TableName: aws.String(db.table),
		Item:      avm,
	}

	if _, err := db.dynamo.PutItem(pii); err != nil {
		return err
	}
	return nil
}

// DeleteQuestion deletes a question in the database by primary key.
func (db *DB) DeleteQuestion(questionID string) error {

	query := &dynamodb.DeleteItemInput{
		TableName: aws.String(db.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(questionID)},
		},
	}

	_, err := db.dynamo.DeleteItem(query)
	if err != nil {
		return err
	}

	return nil
}
