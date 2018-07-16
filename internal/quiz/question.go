package quiz

import (
	"math/rand"
	"reflect"
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

// Question represents a subnet question.
type Question struct {
	ID      string `json:"id"`
	IP      string `json:"ip"`
	Network string `json:"network"`
	Kind    string `json:"kind"`
	TTL     int64  `json:"ttl"`
}

// NewQuestion returns a new randomly generated question struct.
// The fields can be overwritten by parameters.
func NewQuestion(ip, network, kind string) *Question {
	rip, maxSize := randomIP()
	q := &Question{
		ID:      uuid.New().String(),
		IP:      rip,
		Network: randomNetwork(maxSize),
		Kind:    randomQuestionKind(),
		TTL:     time.Now().Add(time.Hour * 8).Unix(),
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

// Solution looks up the function required to provide the solution based upon
// the question kind, transforms the question into a solvable format and solves.
func (q *Question) Solution() string {
	a := newAnswerer(q)
	return solvers[q.Kind](a)
}

// randomIP returns a random private IP as a string with its max cidr.
func randomIP() (string, int) {
	switch r(5) {
	case 0, 1, 2: // In range 10.0.0.0/8
		return addressInRange(10, 0, 255)

	case 3: // In range 172.16.0.0/12
		return addressInRange(172, 16, 31)

	default: // In range 192.168.0.0/16
		return addressInRange(192, 168, 168)
	}
}

// addressInRange returns an address in the required range.
// oct1 = first octet
// oct2 = second octet
// oct2s = second octet start
// oct2f = second octet finish
func addressInRange(oct1, oct2s, oct2f int) (string, int) {
	oct2 := 0
	if oct2s != oct2f {
		oct2 = r(oct2f - oct2s)
	}
	oct2 += oct2s

	maxSize := 12
	if oct1 == 192 {
		maxSize = 16
	}

	o1 := s(oct1)
	o2 := s(oct2)
	o3 := s(r(253) + 1)
	o4 := s(r(253) + 1)

	return o1 + "." + o2 + "." + o3 + "." + o4, maxSize
}

// randomNetwork returns a network in netmask or CIDR notation in the range
// of /12 to /30.
func randomNetwork(maxSize int) string {
	rn := r(minNet - maxSize)
	net := networks[rn]

	switch r(2) {
	case 0:
		return net.netmask
	}
	return net.prefix
}

// randomQuestionKind returns a kind of subnetting question.
func randomQuestionKind() string {
	keys := reflect.ValueOf(solvers).MapKeys()
	rn := r(len(keys))
	return keys[rn].String()
}
