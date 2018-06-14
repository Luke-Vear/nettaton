package nettaton

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Luke-Vear/nettaton/internal/data"
	pf "github.com/Luke-Vear/nettaton/internal/platform"
	"github.com/Luke-Vear/nettaton/internal/quiz"
)

var (
	proxyEvent = `
{
	"body": "{\"test\":\"body\"}",
	"resource": "/{proxy+}",
	"requestContext": {
	  "resourceId": "123456",
	  "apiId": "1234567890",
	  "resourcePath": "/{proxy+}",
	  "httpMethod": "POST",
	  "requestId": "c6af9ac6-7b61-11e6-9a41-93e8deadbeef",
	  "accountId": "123456789012",
	  "identity": {
		"apiKey": null,
		"userArn": null,
		"cognitoAuthenticationType": null,
		"caller": null,
		"userAgent": "Custom User Agent String",
		"user": null,
		"cognitoIdentityPoolId": null,
		"cognitoIdentityId": null,
		"cognitoAuthenticationProvider": null,
		"sourceIp": "127.0.0.1",
		"accountId": null
	  },
	  "stage": "prod"
	},
	"queryStringParameters": {
	  "foo": "bar"
	},
	"headers": {
	  "Via": "1.1 08f323deadbeefa7af34d5feb414ce27.cloudfront.net (CloudFront)",
	  "Accept-Language": "en-US,en;q=0.8",
	  "CloudFront-Is-Desktop-Viewer": "true",
	  "CloudFront-Is-SmartTV-Viewer": "false",
	  "CloudFront-Is-Mobile-Viewer": "false",
	  "X-Forwarded-For": "127.0.0.1, 127.0.0.2",
	  "CloudFront-Viewer-Country": "US",
	  "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
	  "Upgrade-Insecure-Requests": "1",
	  "X-Forwarded-Port": "443",
	  "Host": "1234567890.execute-api.us-east-1.amazonaws.com",
	  "X-Forwarded-Proto": "https",
	  "X-Amz-Cf-Id": "cDehVQoZnx43VYQb9j2-nvCh-9z396Uhbp027Y2JvkCPNLmGJHqlaA==",
	  "CloudFront-Is-Tablet-Viewer": "false",
	  "Cache-Control": "max-age=0",
	  "User-Agent": "Custom User Agent String",
	  "CloudFront-Forwarded-Proto": "https",
	  "Accept-Encoding": "gzip, deflate, sdch"
	},
	"pathParameters": {
	  "proxy": "path/to/resource"
	},
	"httpMethod": "POST",
	"stageVariables": {
	  "baz": "qux"
	},
	"path": "/path/to/resource"
}
`

	storeOK = happyStore{}
	nexusOK = NewNexus(storeOK)

	storeNotOK = unhappyStore{}
	nexusNotOK = NewNexus(storeNotOK)

	storeNotFound = notFoundStore{}
	nexusNotFound = NewNexus(storeNotFound)
)

type happyStore struct{}

func (d happyStore) GetQuestion(string) (*quiz.Question, error) {
	return &quiz.Question{
		ID:      "abc",
		IP:      "10.0.0.0",
		Kind:    "hostsinnet",
		Network: "24",
		TTL:     99999999999,
	}, nil
}
func (d happyStore) UpdateQuestion(*quiz.Question) error    { return nil }
func (d happyStore) DeleteQuestion(questionID string) error { return nil }

type unhappyStore struct{}

func (d unhappyStore) GetQuestion(string) (*quiz.Question, error) {
	return nil, errors.New("update question failed")
}
func (d unhappyStore) UpdateQuestion(*quiz.Question) error {
	return errors.New("update question failed")
}
func (d unhappyStore) DeleteQuestion(questionID string) error {
	return errors.New("update question failed")
}

type notFoundStore struct{}

func (d notFoundStore) GetQuestion(string) (*quiz.Question, error) {
	return nil, data.ErrQuestionNotFound
}
func (d notFoundStore) UpdateQuestion(*quiz.Question) error {
	return data.ErrQuestionNotFound
}
func (d notFoundStore) DeleteQuestion(questionID string) error {
	return data.ErrQuestionNotFound
}

//
func TestNexus_CreateQuestion(t *testing.T) {
	goodRequest := &pf.Request{
		QueryStringParameters: map[string]string{
			"kind": "first",
		},
	}

	badRequest := &pf.Request{
		QueryStringParameters: map[string]string{
			"kind": "foo",
		},
	}

	responseOK, _ := nexusOK.CreateQuestion(goodRequest)
	assert.Equal(t, 201, responseOK.StatusCode)

	var question quiz.Question
	json.Unmarshal([]byte(responseOK.Body), &question)
	assert.Equal(t, "first", question.Kind)

	responseBadKind, _ := nexusOK.CreateQuestion(badRequest)
	assert.Equal(t, 400, responseBadKind.StatusCode)

	responseNotOK, _ := nexusNotOK.CreateQuestion(goodRequest)
	assert.Equal(t, 500, responseNotOK.StatusCode)
}

func TestNexus_ReadQuestion(t *testing.T) {
	goodRequest := &pf.Request{
		PathParameters: map[string]string{
			"id": "abc",
		},
	}

	badRequest := &pf.Request{}

	responseOK, _ := nexusOK.ReadQuestion(goodRequest)
	assert.Equal(t, 200, responseOK.StatusCode)

	responseBadID, _ := nexusOK.ReadQuestion(badRequest)
	assert.Equal(t, 400, responseBadID.StatusCode)

	responseNotFound, _ := nexusNotFound.ReadQuestion(goodRequest)
	assert.Equal(t, 404, responseNotFound.StatusCode)

	responseNotOK, _ := nexusNotOK.ReadQuestion(goodRequest)
	assert.Equal(t, 500, responseNotOK.StatusCode)
}

func TestNexus_AnswerQuestion(t *testing.T) {
	goodRequest := &pf.Request{
		PathParameters: map[string]string{
			"id": "abc",
		},
		Body: `{ "answer": "254" }`,
	}

	badIDRequest := &pf.Request{}

	badBodyRequest := &pf.Request{
		PathParameters: map[string]string{
			"id": "abc",
		},
	}

	responseOK, _ := nexusOK.AnswerQuestion(goodRequest)
	assert.Equal(t, 200, responseOK.StatusCode)

	responseBadID, _ := nexusOK.AnswerQuestion(badIDRequest)
	assert.Equal(t, 400, responseBadID.StatusCode)

	responseBadBody, _ := nexusOK.AnswerQuestion(badBodyRequest)
	assert.Equal(t, 400, responseBadBody.StatusCode)

	responseNotFound, _ := nexusNotFound.AnswerQuestion(goodRequest)
	assert.Equal(t, 404, responseNotFound.StatusCode)

	responseNotOK, _ := nexusNotOK.AnswerQuestion(goodRequest)
	assert.Equal(t, 500, responseNotOK.StatusCode)
}
