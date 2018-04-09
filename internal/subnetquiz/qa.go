package subnetquiz

import (
	"errors"
	"net"
	"regexp"
	"strconv"
)

var (
	// ErrInvalidNetwork unable to parse network.
	ErrInvalidNetwork = errors.New("unable to parse network")

	// Questions maps a string name of a function to an actual function.
	Questions = map[string]func(net.IP, uint) string{
		"first":        First,
		"last":         Last,
		"broadcast":    Broadcast,
		"firstandlast": FirstAndLast,
		"hostsinnet":   HostsInNet,
	}

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
)

// Parse parses an IP address and cidr/netmask into a network address and a cidr.
func Parse(ip, network string) (net.IP, uint, error) {

	// Convert network string to subnet prefix length int.
	cidrStr, err := toCidr(network)
	if err != nil {
		return nil, 0, err
	}
	cidrInt, _ := strconv.Atoi(cidrStr)

	// Convert the strings ip and network into the network ID as type net.IP.
	_, nw, err := net.ParseCIDR(ip + "/" + cidrStr)
	if err != nil {
		return nil, 0, err
	}

	// return IPv4 IP and cidr as uint
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
	return "", ErrInvalidNetwork
}

// First returns the first valid IP address in the range.
func First(nip net.IP, cidr uint) string {

	cpnip := copyNIP(nip)
	cpnip[3]++

	return cpnip.String()
}

// Last returns the last valid IP address in the range.
func Last(nip net.IP, cidr uint) string {

	cpnip := copyNIP(nip)
	hosts := hosts(cidr)

	cpnip[0] = cpnip[0] + byte(hosts/(1<<24))
	cpnip[1] = cpnip[1] + byte(hosts/(1<<16))
	cpnip[2] = cpnip[2] + byte(hosts/(1<<8))
	cpnip[3] = cpnip[3] + byte(hosts)

	return cpnip.String()
}

// Broadcast returns the broadcast address.
func Broadcast(nip net.IP, cidr uint) string {

	bc, _, _ := net.ParseCIDR(Last(nip, cidr) + "/" + strconv.Itoa(int(cidr)))
	bc[len(bc)-1]++

	return bc.String()
}

// FirstAndLast returns the first and last valid IP addresses in the range.
func FirstAndLast(nip net.IP, cidr uint) string {
	return First(nip, cidr) + "-" + Last(nip, cidr)
}

// HostsInNet returns how many valid hosts there are in the subnet.
func HostsInNet(nip net.IP, cidr uint) string {
	return strconv.Itoa(hosts(cidr))
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
