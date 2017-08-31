package cloudplatform

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type mockDB struct{}

func (mdb *mockDB) GetItem(*dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return nil, nil
}

func (mdb *mockDB) DeleteItem(*dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	return nil, nil
}

func (mdb *mockDB) PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

func TestNewUser(t *testing.T) {
	name := "somethingOriginal"
	user := NewUser(name)
	if user.ID != name {
		t.Errorf("username - actual: %v, expected: %v", user.ID, name)
	}

	if len(user.Marks) == 0 {
		t.Errorf("usermarks should be initialized - len: %v, actual: %v, ", len(user.Marks), user.Marks)
	}

	for _, val := range user.Marks {
		if val.Attempts != 0 || val.Correct != 0 {
			t.Errorf("marks should be zero - actual: %v, %v", val.Attempts, val.Correct)
		}
	}
}
