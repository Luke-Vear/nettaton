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

// Question represents a subnet question.
type Question struct {
	ID      string `json:"id"`
	IP      string `json:"ip"`
	Network string `json:"network"`
	Kind    string `json:"kind"`
}

// NewQuestion returns a new randomly generated question struct.
// The fields can be overwritten by parameters.
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

// Solution looks up the function required to provide the solution based upon
// the question kind, transforms the question into a solvable format and solves.
func (q *Question) Solution() string {
	a := newAnswerer(q)
	return solvers[q.Kind](a)
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

// addressInRange returns an address in the required range.
// oct1 = first octet
// oct2 = second octet
// oct2s = second octet start
// oct2f = second octet finish
func addressInRange(oct1, oct2s, oct2f int) string {
	oct2 := 0
	if oct2s != oct2f {
		oct2 = r(oct2f - oct2s)
	}
	oct2 += oct2s

	o1 := s(oct1)
	o2 := s(oct2)
	o3 := s(r(253) + 1)
	o4 := s(r(253) + 1)

	return o1 + "." + o2 + "." + o3 + "." + o4
}

// randomNetwork returns a network in netmask or CIDR notation in the range
// of /12 to /30.
func randomNetwork() string {
	ln := len(solvers)
	rn := r(ln)
	net := networks[rn]

	switch r(2) {
	case 0:
		return net.netmask
	}
	return net.prefix
}

// randomQuestionKind returns a kind of subnetting question.
func randomQuestionKind() string {
	var qks []string
	for key := range solvers {
		qks = append(qks, key)
	}

	ls := len(solvers)
	rn := r(ls)
	return qks[rn]
}
