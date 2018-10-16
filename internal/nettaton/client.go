package nettaton

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	serverFQDN string
	ct         string
	http       *http.Client
}

func NewClient(serverFQDN string) *Client {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}
	return &Client{
		serverFQDN: serverFQDN,
		ct:         "application/json",
		http:       httpClient,
	}
}

func (c *Client) questionEndpoint() string {
	return "https://" + c.serverFQDN + "/question"
}

// CreateQuestion ...
func (c *Client) CreateQuestion() (string, error) {
	resp, err := c.http.Post(c.questionEndpoint(), c.ct, nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("status code: %v, error: %v", resp.StatusCode, err)
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return bbToJSON(bb)
}

func (c *Client) ReadQuestion(uuid uuid.UUID) (string, error) {
	endpoint := c.questionEndpoint() + "/" + uuid.String()

	resp, err := c.http.Get(endpoint)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %v, error: %v", resp.StatusCode, err)
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return bbToJSON(bb)
}

func (c *Client) answerEndpoint(id string) string {
	return c.questionEndpoint() + "/" + id + "/answer"
}

func (c *Client) AnswerQuestion(uuid uuid.UUID, answer string) (string, error) {
	endpoint := c.answerEndpoint(uuid.String())

	resp, err := c.http.Post(endpoint, c.ct, reader(answer))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %v, error: %v", resp.StatusCode, err)
	}

	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return bbToJSON(bb)
}

func reader(answer string) io.Reader {
	json, _ := json.Marshal(struct {
		Answer string `json:"answer"`
	}{
		Answer: answer,
	})
	return bytes.NewReader(json)
}

func bbToJSON(bb []byte) (string, error) {
	var jsonBB bytes.Buffer

	err := json.Indent(&jsonBB, bb, "", "  ")
	if err != nil {
		return "", err
	}

	return jsonBB.String(), nil
}
