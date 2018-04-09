package subnetquiz

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	// Too verbose, I use lots.
	s, r = strconv.Itoa, rand.Intn
)

// Need to seed for random question generation below.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandomIP returns a random IP in the range 10.0.0.0/8 as a string.
// TODO: more real internal IP ranges.
func RandomIP() string {
	return "10" + "." + s(r(253)+1) + "." + s(r(253)+1) + "." + s(r(253)+1)
}

// RandomNetwork returns a network in netmask or CIDR notation in the range
// of /12 to /30.
func RandomNetwork() string {
	switch r(2) {
	case 0:
		return Networks[r(len(Networks))].netmask
	}
	return Networks[r(len(Networks))].prefix
}

// RandomQuestionKind returns a kind of subnetting question.
func RandomQuestionKind() string {

	var qk []string
	for key := range Questions {
		qk = append(qk, key)
	}

	return qk[r(len(Questions))]
}
