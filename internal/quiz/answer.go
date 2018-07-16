package quiz

import (
	"errors"
	"net"
	"strconv"
)

var (
	// ErrInvalidNetwork unable to parse network.
	ErrInvalidNetwork = errors.New("unable to parse network")
	// ErrInvalidQuestionKind invalid question kind.
	ErrInvalidQuestionKind = errors.New("invalid question kind")

	// solvers maps a question kind to the function that can solve it.
	solvers = map[string]func(*answerer) string{
		"first":        (*answerer).first,
		"last":         (*answerer).last,
		"broadcast":    (*answerer).broadcast,
		"firstandlast": (*answerer).firstAndLast,
		"hostsinnet":   (*answerer).hostsInNet,
	}

	// networks in netmask or CIDR notation in the range of /12 to /30.
	networks = []struct {
		netmask string
		prefix  string
	}{
		{netmask: "255.255.255.252", prefix: "30"},
		{netmask: "255.255.255.248", prefix: "29"},
		{netmask: "255.255.255.240", prefix: "28"},
		{netmask: "255.255.255.224", prefix: "27"},
		{netmask: "255.255.255.192", prefix: "26"},
		{netmask: "255.255.255.128", prefix: "25"},
		{netmask: "255.255.255.0", prefix: "24"},
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

	maxNet = 12
	minNet = 30
)

// answerer can answer questions about a subnet. It is created from a Question
// by parsing the question data into a format that can be used for calculation.
type answerer struct {
	nip  net.IP
	cidr uint
}

// newAnswerer parses an IP address and cidr/netmask into an answerer.
func newAnswerer(q *Question) *answerer {
	// Convert network string to subnet prefix length int.
	cidrStr := toCidr(q.Network)
	cidrInt, _ := strconv.Atoi(cidrStr)

	// Convert the strings ip and network into the network ID as type net.IP.
	_, nw, _ := net.ParseCIDR(q.IP + "/" + cidrStr)

	// return IPv4 IP and cidr as uint
	return &answerer{
		nip:  nw.IP.To4(),
		cidr: uint(cidrInt),
	}
}

// toCidr transforms netmask to cidr, or returns the cidr if in correct range.
func toCidr(n string) string {
	for _, v := range networks {
		if v.netmask == n {
			return v.prefix
		}
	}
	return n
}

// first returns the first valid IP address in the range.
func (a *answerer) first() string {
	cpnip := copyNIP(a.nip)
	cpnip[3]++

	return cpnip.String()
}

// last returns the last valid IP address in the range.
func (a *answerer) last() string {
	cpnip := copyNIP(a.nip)
	hosts := hosts(a.cidr)

	cpnip[0] = cpnip[0] + byte(hosts/(1<<24))
	cpnip[1] = cpnip[1] + byte(hosts/(1<<16))
	cpnip[2] = cpnip[2] + byte(hosts/(1<<8))
	cpnip[3] = cpnip[3] + byte(hosts)

	return cpnip.String()
}

// broadcast returns the broadcast address.
func (a *answerer) broadcast() string {
	bc, _, _ := net.ParseCIDR(a.last() + "/" + strconv.Itoa(int(a.cidr)))
	bc[len(bc)-1]++

	return bc.String()
}

// firstAndLast returns the first and last valid IP addresses in the range.
func (a *answerer) firstAndLast() string {
	return a.first() + "-" + a.last()
}

// hostsInNet returns how many valid hosts there are in the subnet.
func (a *answerer) hostsInNet() string {
	return strconv.Itoa(hosts(a.cidr))
}

// copyNIP returns a copy of the net.IP to prevent global state mutation.
func copyNIP(nip net.IP) net.IP {
	cpnip := make(net.IP, len(nip))
	copy(cpnip, nip)
	return cpnip
}

// hosts returns the amount of valid hosts in the subnet as an int.
func hosts(cidr uint) int {
	return 2<<(31-cidr) - 2
}

// ValidQuestionKind returns true if the kind is valid.
func ValidQuestionKind(kind string) bool {
	_, ok := solvers[kind]
	return ok
}
