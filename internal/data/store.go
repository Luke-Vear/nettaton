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
	table  *string
}

// NewStore returns a instantiated Store.
func NewStore(table string, dynamo *dynamodb.DynamoDB) *Store {
	return &Store{
		table:  aws.String(table),
		dynamo: dynamo,
	}
}

// GetQuestion retrieves a question in the database by primary key.
func (ds *Store) GetQuestion(questionID string) (*quiz.Question, error) {

	qid := aws.String(questionID)

	gii := &dynamodb.GetItemInput{
		TableName: ds.table,
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: qid}},
	}

	result, err := ds.dynamo.GetItem(gii)
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
func (ds *Store) UpdateQuestion(qq *quiz.Question) error {

	avm, err := dynamodbattribute.MarshalMap(qq)
	if err != nil {
		return err
	}

	pii := &dynamodb.PutItemInput{
		TableName: ds.table,
		Item:      avm,
	}

	if _, err := ds.dynamo.PutItem(pii); err != nil {
		return err
	}
	return nil
}

// DeleteQuestion deletes a question in the database by primary key.
func (ds *Store) DeleteQuestion(questionID string) error {

	qid := aws.String(questionID)

	dii := &dynamodb.DeleteItemInput{
		TableName: ds.table,
		Key:       map[string]*dynamodb.AttributeValue{"id": {S: qid}},
	}

	_, err := ds.dynamo.DeleteItem(dii)
	if err != nil {
		return err
	}

	return nil
}
