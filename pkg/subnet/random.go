package subnet

import (
	"math/rand"
	"strconv"
)

var (
	// Networks in netmask or CIDR notation in the range of /12 to /30.
	Networks = []struct {
		netmask string
		prefix  string
	}{
		{netmask: "255.255.255.252", prefix: "30"}, // 2 valid hosts
		{netmask: "255.255.255.248", prefix: "29"},
		{netmask: "255.255.255.240", prefix: "28"},
		{netmask: "255.255.255.224", prefix: "27"},
		{netmask: "255.255.255.192", prefix: "26"},
		{netmask: "255.255.255.128", prefix: "25"},
		{netmask: "255.255.255.0", prefix: "24"}, // 254 valid hosts
		{netmask: "255.255.254.0", prefix: "23"},
		{netmask: "255.255.252.0", prefix: "22"},
		{netmask: "255.255.248.0", prefix: "21"},
		{netmask: "255.255.240.0", prefix: "20"},
		{netmask: "255.255.224.0", prefix: "19"},
		{netmask: "255.255.192.0", prefix: "18"},
		{netmask: "255.255.128.0", prefix: "17"},
		{netmask: "255.255.0.0", prefix: "16"},
		{netmask: "255.254.0.0", prefix: "15"},
		{netmask: "255.252.0.0", prefix: "14"},
		{netmask: "255.248.0.0", prefix: "13"},
		{netmask: "255.240.0.0", prefix: "12"},
	}

	// Too verbose, I use lots.
	s, r = strconv.Itoa, rand.Intn
)

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

	var questionKinds []string
	for key := range QuestionFuncMap {
		questionKinds = append(questionKinds, key)
	}

	return questionKinds[r(len(QuestionFuncMap))]
}
