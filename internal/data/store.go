package data

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

// Store represents a physical datastore.
type Store struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

// NewStore returns a instantiated Store.
func NewStore(table string, dynamo *dynamodb.DynamoDB) *Store {
	return &Store{
		table:  table,
		dynamo: dynamo,
	}
}

// GetQuestion retrieves a question in the database by primary key.
func (ss *Store) GetQuestion(questionID string) (*quiz.Question, error) {

	query := &dynamodb.GetItemInput{
		TableName: aws.String(ss.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(questionID)},
		},
	}

	result, err := ss.dynamo.GetItem(query)
	if err != nil {
		return nil, err
	}

	var qq *quiz.Question
	if err := dynamodbattribute.UnmarshalMap(result.Item, &qq); err != nil {
		return nil, err
	}

	if len(qq.IP) == 0 {
		return nil, ErrQuestionNotFound
	}

	return qq, nil
}

// UpdateQuestion creates or overwrites a question in the database by primary key.
func (ss *Store) UpdateQuestion(qq *quiz.Question) error {

	avm, err := dynamodbattribute.MarshalMap(qq)
	if err != nil {
		return err
	}

	pii := &dynamodb.PutItemInput{
		TableName: aws.String(ss.table),
		Item:      avm,
	}

	if _, err := ss.dynamo.PutItem(pii); err != nil {
		return err
	}
	return nil
}

// DeleteQuestion deletes a question in the database by primary key.
func (ss *Store) DeleteQuestion(questionID string) error {

	query := &dynamodb.DeleteItemInput{
		TableName: aws.String(ss.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(questionID)},
		},
	}

	_, err := ss.dynamo.DeleteItem(query)
	if err != nil {
		return err
	}

	return nil
}
