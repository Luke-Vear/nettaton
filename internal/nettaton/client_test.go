package nettaton

import (
	"testing"
)

var (
	question = `{"id":"6d337539-2e3f-406e-b4f1-982b6053ea47","ip":"192.168.163.112","network":"25","kind":"first","ttl":1532653658}`
	answer   = `{"correct":true}`
)

// {"id":"6d337539-2e3f-406e-b4f1-982b6053ea47","ip":"192.168.163.112","network":"25","kind":"first","ttl":1532653658}
func TestClient_CreateQuestion(t *testing.T) {
	// ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello, client")
	// }))
	// defer ts.Close()

	// res, err := http.Get(ts.URL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

// {"id":"6d337539-2e3f-406e-b4f1-982b6053ea47","ip":"192.168.163.112","network":"25","kind":"first","ttl":1532653658}

func TestClient_ReadQuestion(t *testing.T) {}

// {"correct":true}
func TestClient_AnswerQuestion(t *testing.T) {}
