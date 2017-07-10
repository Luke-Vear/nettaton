package subnet

import (
	"fmt"
	"math/rand"
	"net"
	"regexp"
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

	// QuestionFuncMap maps a string name of a function to an actual function.
	QuestionFuncMap = map[string]func(net.IP, uint) string{
		"first":        First,
		"last":         Last,
		"broadcast":    Broadcast,
		"firstandlast": FirstAndLast,
		"hostsinnet":   HostsInNet,
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

// Parse parses a
func Parse(ip, network string) (net.IP, uint, error) {

	cidrStr, err := toCidr(network)
	if err != nil {
		return nil, 0, err
	}

	cidrInt, err := strconv.Atoi(cidrStr)
	if err != nil {
		return nil, 0, err
	}

	_, nw, err := net.ParseCIDR(ip + "/" + cidrStr)
	if err != nil {
		return nil, 0, err
	}

	return nw.IP.To4(), uint(cidrInt), nil
}

// toCidr returns the cidr if in correct range, or transforms netmask to cidr.
func toCidr(n string) (string, error) {
	if match, _ := regexp.MatchString("^(1[2-9]|2[0-9]|30)$", n); match {
		return n, nil
	}
	for _, v := range Networks {
		if v.netmask == n {
			return v.prefix, nil
		}
	}
	return "", fmt.Errorf("Unable to parse network %v", n)
}

// First returns.
func First(nip net.IP, cidr uint) string {

	nip[3]++

	return nip.String()
}

// Last returns.
func Last(nip net.IP, cidr uint) string {

	hosts := 2<<(31-cidr) - 2

	nip[0] = nip[0] + byte(hosts/(1<<24))
	nip[1] = nip[1] + byte(hosts/(1<<16))
	nip[2] = nip[2] + byte(hosts/(1<<8))
	nip[3] = nip[3] + byte(hosts)

	return nip.String()
}

// Broadcast returns.
func Broadcast(nip net.IP, cidr uint) string {

	bc := net.ParseIP(Last(nip, cidr) + "/" + string(cidr))
	bc[3]++

	return bc.String()
}

// FirstAndLast returns.
func FirstAndLast(nip net.IP, cidr uint) string {
	return First(nip, cidr) + "-" + Last(nip, cidr)
}

// HostsInNet returns.
func HostsInNet(nip net.IP, cidr uint) string {
	return strconv.Itoa(1<<(32-cidr) - 2)
}
