package cloudplatform

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Luke-Vear/nettaton/pkg/subnetquiz"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var ErrForTest = errors.New("error for test")

func init() {
	db = &mockDB{}
	table = "noTable"
}

type mockDB struct{}

func (mdb *mockDB) GetItem(gii *dynamodb.GetItemInput) (gio *dynamodb.GetItemOutput, err error) {
	errorName := &dynamodb.AttributeValue{S: aws.String("errorName")}
	if reflect.DeepEqual(gii.Key["id"], errorName) {
		return nil, ErrForTest
	}

	notFoundName := &dynamodb.AttributeValue{S: aws.String("notFoundName")}
	if reflect.DeepEqual(gii.Key["id"], notFoundName) {
		return &dynamodb.GetItemOutput{
			Item: map[string]*dynamodb.AttributeValue{
				"id":     notFoundName,
				"status": &dynamodb.AttributeValue{S: aws.String("")},
			},
		}, nil
	}

	allFineName := &dynamodb.AttributeValue{S: aws.String("anyNameWillDo")}
	gio = &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"id":     allFineName,
			"status": &dynamodb.AttributeValue{S: aws.String("justNotEmpty")},
		},
	}
	return gio, nil
}

func (mdb *mockDB) PutItem(pii *dynamodb.PutItemInput) (pio *dynamodb.PutItemOutput, err error) {
	if *pii.TableName == "mockError" {
		return nil, ErrForTest
	}
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

func TestUser_GradeAnswer(t *testing.T) {
	user := NewUser("somename")
	qk := subnetquiz.RandomQuestionKind()

	user.GradeAnswer(false, qk)
	oneWrong := &Marks{1, 0}
	if !reflect.DeepEqual(user.Marks[qk], oneWrong) {
		t.Errorf("attempt should be one, correct should be zero: %v", user.Marks[qk])
	}

	user.GradeAnswer(true, qk)
	twoAttemptOneRight := &Marks{2, 1}
	if !reflect.DeepEqual(user.Marks[qk], twoAttemptOneRight) {
		t.Errorf("attempt should be two, correct should be one: %v", user.Marks[qk])
	}
}

func TestUser_ListMarks(t *testing.T) {
	user := NewUser("somename")
	qk := subnetquiz.RandomQuestionKind()

	user.GradeAnswer(false, qk)
	oneWrong := &Marks{1, 0}
	if !reflect.DeepEqual(user.ListMarks()[qk], oneWrong) {
		t.Errorf("attempt should be one, correct should be zero: %v", user.Marks[qk])
	}
}

func TestUser_isNotFound(t *testing.T) {
	user := NewUser("somename")
	if !user.isNotFound() {
		t.Error("user should have empty found field")
	}

	user.Status = "anything"
	if user.isNotFound() {
		t.Error("user should not have empty found field")
	}
}

func TestUser_Login(t *testing.T) {
	user := NewUser("somename")
	user.HashedPassword = "$2a$10$YZtNOffXRIekdOzILPokJuaX1Yn5qIi2bEY1kbPWAcTvdHl77dqca" //testPassword123

	_, err := user.Login("testPassword123")
	if err != nil {
		t.Errorf("actualErr: %v, expected: nil", err)
	}

	_, err = user.Login("notCorrect")
	expectedErr := "crypto/bcrypt: hashedPassword is not the hash of the given password"
	if err.Error() != expectedErr {
		t.Errorf("actualErr: %v, expectedErr: %v", err, expectedErr)
	}
}

func TestUser_Create(t *testing.T) {
	u := NewUser("somename")
	err := u.Create()
	if err != ErrUserAlreadyExists {
		t.Errorf("actualErr: %v, expectErr: %v", err, ErrUserAlreadyExists)
	}

	eu := NewUser("errorName")
	expectErr := "error for test"
	if err := eu.Read(); err.Error() != expectErr {
		t.Errorf("actualErr: %v, expectErr: %v", err, expectErr)
	}

	nfu := NewUser("notFoundName")
	nfu.ClearTextPassword = "testPassword123"
	_ = nfu.Create()
	expectStatus := "new"
	if nfu.Status != expectStatus {
		t.Errorf("actualStatus: %v, expectStatus: %v", nfu.Status, expectStatus)
	}
	if len(nfu.HashedPassword) == 0 {
		t.Fatalf("hash password length 0, password not hashed")
	}
	if string(nfu.HashedPassword[0]) != "$" {
		t.Errorf("hashed password incorrect format: %v", nfu.HashedPassword)
	}
}

func TestUser_Read(t *testing.T) {
	eu := NewUser("errorName")
	expectErr := "error for test"
	if err := eu.Read(); err.Error() != expectErr {
		t.Errorf("actualErr: %v, expectErr: %v", err, expectErr)
	}

	nfu := NewUser("notFoundName")
	if err := nfu.Read(); err != ErrUserNotFoundInDatabase {
		t.Errorf("actualErr: %v, expectErr: %v", err, ErrUserNotFoundInDatabase)
	}

	u := NewUser("pat")
	err := u.Read()
	if err != nil {
		t.Errorf("actualErr: %v, expect nil err", err)
	}
	expectStatus := "justNotEmpty"
	if u.Status != expectStatus {
		t.Errorf("actualStatus: %v, expectStatus: %v", u.Status, expectStatus)
	}
}

func TestUser_Update(t *testing.T) {
	u := NewUser("nameNotImportant")
	u.ClearTextPassword = "testPassword123"
	_ = u.Update()
	if u.ClearTextPassword != "" {
		t.Errorf("not clearing u.ClearTextPassword")
	}

	table = "mockError"
	if err := u.Update(); err != ErrForTest {
		t.Errorf("actualErr: %v, expectErr: %v", err, ErrForTest)
	}
}
