package nettaton

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	pf "github.com/Luke-Vear/nettaton/internal/platform"
	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/Luke-Vear/nettaton/internal/state"
)

var (
	storeOK   = happyStore{}
	gatewayOK = state.NewGateway(storeOK)
	nexusOK   = NewNexus(gatewayOK)

	storeNotOK   = unhappyStore{}
	gatewayNotOK = state.NewGateway(storeNotOK)
	nexusNotOK   = NewNexus(gatewayNotOK)
)

type happyStore struct{}

func (d happyStore) GetQuestion(string) (*quiz.Question, error) { return nil, nil }
func (d happyStore) UpdateQuestion(*quiz.Question) error        { return nil }
func (d happyStore) DeleteQuestion(questionID string) error     { return nil }

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

	responseBadRequestOK, _ := nexusOK.CreateQuestion(badRequest)
	assert.Equal(t, 400, responseBadRequestOK.StatusCode)

	responseNotOK, _ := nexusNotOK.CreateQuestion(goodRequest)
	assert.Equal(t, 500, responseNotOK.StatusCode)
}

func TestNexus_ReadQuestion(t *testing.T) {}

func TestNexus_AnswerQuestion(t *testing.T) {

}

var proxyEvent = `
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
