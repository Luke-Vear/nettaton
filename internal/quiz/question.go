package quiz

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var (
	// Too verbose, I use lots.
	s, r = strconv.Itoa, rand.Intn
)

// Need to seed for random question generation below.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Question ...
type Question struct {
	ID      string `json:"id"`
	IP      string `json:"ip"`
	Network string `json:"network"`
	Kind    string `json:"kind"`
}

// NewQuestion ...
func NewQuestion(ip, network, kind string) *Question {
	q := &Question{
		ID:      uuid.New().String(),
		IP:      randomIP(),
		Network: randomNetwork(),
		Kind:    randomQuestionKind(),
	}
	if len(ip) > 0 {
		q.IP = ip
	}
	if len(network) > 0 {
		q.Network = network
	}
	if len(kind) > 0 {
		q.Kind = kind
	}
	return q
}

// randomIP returns a random private IP as a string.
func randomIP() string {
	switch r(3) {

	case 0: // In range 10.0.0.0/8
		return addressInRange(10, 0, 255)

	case 1: // In range 172.16.0.0/12
		return addressInRange(172, 16, 31)

	default: // In range 192.168.0.0/16
		return addressInRange(192, 168, 168)
	}

}

func addressInRange(fo, sos, sof int) string {
	so := 0
	if sos != sof {
		so = r(sof - sos)
	}
	so += sos

	return s(fo) + "." + s(so) + "." + s(r(253)+1) + "." + s(r(253)+1)
}

// randomNetwork returns a network in netmask or CIDR notation in the range
// of /12 to /30.
func randomNetwork() string {
	switch r(2) {
	case 0:
		return networks[r(len(networks))].netmask
	}
	return networks[r(len(networks))].prefix
}

// randomQuestionKind returns a kind of subnetting question.
func randomQuestionKind() string {

	var qk []string
	for key := range questions {
		qk = append(qk, key)
	}

	return qk[r(len(questions))]
}
