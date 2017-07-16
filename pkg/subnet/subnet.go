package subnet

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

var (
	// QuestionFuncMap maps a string name of a function to an actual function.
	QuestionFuncMap = map[string]func(net.IP, uint) string{
		"first":        First,
		"last":         Last,
		"broadcast":    Broadcast,
		"firstandlast": FirstAndLast,
		"hostsinnet":   HostsInNet,
	}
)

// Parse parses an IP address and cidr/netmask into a network address and a cidr.
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

// First returns the first valid IP address in the range.
func First(nip net.IP, cidr uint) string {

	cpnip := cpnip(nip)
	cpnip[3]++

	return cpnip.String()
}

// Last returns the last valid IP address in the range.
func Last(nip net.IP, cidr uint) string {

	cpnip := cpnip(nip)
	hosts := hosts(cidr)

	cpnip[0] = cpnip[0] + byte(hosts/(1<<24))
	cpnip[1] = cpnip[1] + byte(hosts/(1<<16))
	cpnip[2] = cpnip[2] + byte(hosts/(1<<8))
	cpnip[3] = cpnip[3] + byte(hosts)

	return cpnip.String()
}

// Broadcast returns the broadcast address.
func Broadcast(nip net.IP, cidr uint) string {

	bc := net.ParseIP(Last(nip, cidr) + "/" + string(cidr))
	bc[3]++

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

// cpnip returns a copy of the net.IP to prevent global state mutation.
func cpnip(nip net.IP) net.IP {
	cpnip := make(net.IP, len(nip))
	copy(cpnip, nip)
	return cpnip
}

// hosts returns the amount of valid hosts in the subnet as an int.
func hosts(cidr uint) int {
	return 2<<(31-cidr) - 2
}
