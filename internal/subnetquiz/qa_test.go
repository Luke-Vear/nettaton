package subnetquiz

import (
	"net"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	tt := []struct {
		inputIP    string
		inputNet   string
		netExpect  net.IP
		cidrExpect uint
		errExpect  error
	}{
		{
			inputIP:    "",
			inputNet:   "31",
			netExpect:  nil,
			cidrExpect: 0,
			errExpect:  ErrInvalidNetwork,
		},
		{
			inputIP:    "",
			inputNet:   "255.255.255.254",
			netExpect:  nil,
			cidrExpect: 0,
			errExpect:  ErrInvalidNetwork,
		},
		{
			inputIP:    "10.72.243.106",
			inputNet:   "30",
			netExpect:  net.IP{10, 72, 243, 104},
			cidrExpect: 30,
			errExpect:  nil,
		},
		{
			inputIP:    "10.16.9.99",
			inputNet:   "255.240.0.0",
			netExpect:  net.IP{10, 16, 0, 0},
			cidrExpect: 12,
			errExpect:  nil,
		},
		{
			inputIP:    "",
			inputNet:   "11",
			netExpect:  nil,
			cidrExpect: 0,
			errExpect:  ErrInvalidNetwork,
		},
		{
			inputIP:    "",
			inputNet:   "255.224.0.0",
			netExpect:  nil,
			cidrExpect: 0,
			errExpect:  ErrInvalidNetwork,
		},
		{
			inputIP:    "thisWillError",
			inputNet:   "255.240.0.0",
			netExpect:  nil,
			cidrExpect: 0,
			errExpect:  error(&net.ParseError{Type: "CIDR address", Text: "thisWillError/12"}),
		},
	}

	for _, tc := range tt {
		netActual, cidrActual, errActual := Parse(tc.inputIP, tc.inputNet)
		if !reflect.DeepEqual(netActual, tc.netExpect) || cidrActual != tc.cidrExpect || !reflect.DeepEqual(errActual, tc.errExpect) {
			t.Errorf("\nnetActual: %v\nnetExpect: %v\ncidrActual: %v\ncidrExpect: %v\nerrActual: %v\nerrExpect: %v\n",
				netActual, tc.netExpect, cidrActual, tc.cidrExpect, errActual, tc.errExpect)
		}
	}
}

func TestToCidr(t *testing.T) {

	tt := []struct {
		input     string
		errExpect error
		strExpect string
	}{
		{
			input:     "31",
			errExpect: ErrInvalidNetwork,
			strExpect: "",
		},
		{
			input:     "255.255.255.254",
			errExpect: ErrInvalidNetwork,
			strExpect: "",
		},
		{
			input:     "30",
			errExpect: nil,
			strExpect: "30",
		},
		{
			input:     "255.255.255.252",
			errExpect: nil,
			strExpect: "30",
		},
		{
			input:     "12",
			errExpect: nil,
			strExpect: "12",
		},
		{
			input:     "255.240.0.0",
			errExpect: nil,
			strExpect: "12",
		},
		{
			input:     "11",
			errExpect: ErrInvalidNetwork,
			strExpect: "",
		},
		{
			input:     "255.224.0.0",
			errExpect: ErrInvalidNetwork,
			strExpect: "",
		},
	}

	for _, tc := range tt {
		if actualStr, actualErr := toCidr(tc.input); actualStr != tc.strExpect || actualErr != tc.errExpect {
			t.Errorf("\ntc: %v\nactualStr: %v, actualErr: %v\n", tc, actualStr, actualErr)
		}
	}
}

func TestFirst(t *testing.T) {

	tt := []struct {
		inputNetIP net.IP
		inputCidr  uint
		expect     string
	}{
		{
			inputNetIP: net.IP{10, 72, 243, 104},
			inputCidr:  30,
			expect:     "10.72.243.105",
		},
		{
			inputNetIP: net.IP{10, 28, 244, 152},
			inputCidr:  29,
			expect:     "10.28.244.153",
		},
		{
			inputNetIP: net.IP{10, 217, 75, 0},
			inputCidr:  24,
			expect:     "10.217.75.1",
		},
		{
			inputNetIP: net.IP{10, 16, 0, 0},
			inputCidr:  12,
			expect:     "10.16.0.1",
		},
	}

	for _, tc := range tt {
		if actual := First(tc.inputNetIP, tc.inputCidr); actual != tc.expect {
			t.Errorf("\nactual: %v\nexpected: %v\nip: %v, cidr: %v", actual, tc.expect, tc.inputNetIP, tc.inputCidr)
		}
	}
}

func TestLast(t *testing.T) {

	tt := []struct {
		inputNetIP net.IP
		inputCidr  uint
		expect     string
	}{
		{
			inputNetIP: net.IP{10, 72, 243, 104},
			inputCidr:  30,
			expect:     "10.72.243.106",
		},
		{
			inputNetIP: net.IP{10, 28, 244, 152},
			inputCidr:  29,
			expect:     "10.28.244.158",
		},
		{
			inputNetIP: net.IP{10, 217, 75, 0},
			inputCidr:  24,
			expect:     "10.217.75.254",
		},
		{
			inputNetIP: net.IP{10, 16, 0, 0},
			inputCidr:  12,
			expect:     "10.31.255.254",
		},
	}

	for _, tc := range tt {
		if actual := Last(tc.inputNetIP, tc.inputCidr); actual != tc.expect {
			t.Errorf("\nactual: %v\nexpected: %v\nip: %v, cidr: %v", actual, tc.expect, tc.inputNetIP, tc.inputCidr)
		}
	}
}

func TestBroadcast(t *testing.T) {

	tt := []struct {
		inputNetIP net.IP
		inputCidr  uint
		expect     string
	}{
		{
			inputNetIP: net.IP{10, 72, 243, 104},
			inputCidr:  30,
			expect:     "10.72.243.107",
		},
		{
			inputNetIP: net.IP{10, 28, 244, 152},
			inputCidr:  29,
			expect:     "10.28.244.159",
		},
		{
			inputNetIP: net.IP{10, 217, 75, 0},
			inputCidr:  24,
			expect:     "10.217.75.255",
		},
		{
			inputNetIP: net.IP{10, 16, 0, 0},
			inputCidr:  12,
			expect:     "10.31.255.255",
		},
	}

	for _, tc := range tt {
		if actual := Broadcast(tc.inputNetIP, tc.inputCidr); actual != tc.expect {
			t.Errorf("\nactual: %v\nexpected: %v\nip: %v, cidr: %v", actual, tc.expect, tc.inputNetIP, tc.inputCidr)
		}
	}
}

func TestFirstAndLast(t *testing.T) {

	tt := []struct {
		inputNetIP net.IP
		inputCidr  uint
		expect     string
	}{
		{
			inputNetIP: net.IP{10, 72, 243, 104},
			inputCidr:  30,
			expect:     "10.72.243.105-10.72.243.106",
		},
		{
			inputNetIP: net.IP{10, 28, 244, 152},
			inputCidr:  29,
			expect:     "10.28.244.153-10.28.244.158",
		},
		{
			inputNetIP: net.IP{10, 217, 75, 0},
			inputCidr:  24,
			expect:     "10.217.75.1-10.217.75.254",
		},
		{
			inputNetIP: net.IP{10, 16, 0, 0},
			inputCidr:  12,
			expect:     "10.16.0.1-10.31.255.254",
		},
	}

	for _, tc := range tt {
		if actual := FirstAndLast(tc.inputNetIP, tc.inputCidr); actual != tc.expect {
			t.Errorf("\nactual: %v\nexpected: %v\nip: %v, cidr: %v", actual, tc.expect, tc.inputNetIP, tc.inputCidr)
		}
	}
}

func TestHostsInNet(t *testing.T) {

	tt := []struct {
		inputNetIP net.IP
		inputCidr  uint
		expect     string
	}{
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  12,
			expect:     "1048574",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  13,
			expect:     "524286",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  14,
			expect:     "262142",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  15,
			expect:     "131070",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  16,
			expect:     "65534",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  17,
			expect:     "32766",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  18,
			expect:     "16382",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  19,
			expect:     "8190",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  20,
			expect:     "4094",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  21,
			expect:     "2046",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  22,
			expect:     "1022",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  23,
			expect:     "510",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  24,
			expect:     "254",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  25,
			expect:     "126",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  26,
			expect:     "62",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  27,
			expect:     "30",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  28,
			expect:     "14",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  29,
			expect:     "6",
		},
		{
			inputNetIP: net.IP{10, 10, 10, 10},
			inputCidr:  30,
			expect:     "2",
		},
	}

	for _, tc := range tt {
		if actual := HostsInNet(tc.inputNetIP, tc.inputCidr); actual != tc.expect {
			t.Errorf("\nactual: %v\nexpected: %v\nip: %v, cidr: %v", actual, tc.expect, tc.inputNetIP, tc.inputCidr)
		}
	}

}

func TestCopyNIP(t *testing.T) {

	ipToBeCopied, ipToBeCompared := net.IP{10, 10, 10, 10}, net.IP{10, 10, 10, 10}

	// Copy IP and mutate original.
	copiedIP := copyNIP(ipToBeCopied)
	ipToBeCopied[len(ipToBeCopied)-1]++

	if copiedIP.Equal(ipToBeCopied) {
		t.Error("IP mutated, not copied")
	}

	if !copiedIP.Equal(ipToBeCompared) {
		t.Error("IP not copied correctly")
	}
}

func TestHosts(t *testing.T) {

	tt := []struct {
		input  uint
		expect int
	}{
		{
			input:  12,
			expect: 1048574,
		},
		{
			input:  13,
			expect: 524286,
		},
		{
			input:  14,
			expect: 262142,
		},
		{
			input:  15,
			expect: 131070,
		},
		{
			input:  16,
			expect: 65534,
		},
		{
			input:  17,
			expect: 32766,
		},
		{
			input:  18,
			expect: 16382,
		},
		{
			input:  19,
			expect: 8190,
		},
		{
			input:  20,
			expect: 4094,
		},
		{
			input:  21,
			expect: 2046,
		},
		{
			input:  22,
			expect: 1022,
		},
		{
			input:  23,
			expect: 510,
		},
		{
			input:  24,
			expect: 254,
		},
		{
			input:  25,
			expect: 126,
		},
		{
			input:  26,
			expect: 62,
		},
		{
			input:  27,
			expect: 30,
		},
		{
			input:  28,
			expect: 14,
		},
		{
			input:  29,
			expect: 6,
		},
		{
			input:  30,
			expect: 2,
		},
	}

	for _, tc := range tt {
		if actual := hosts(tc.input); actual != tc.expect {
			t.Errorf("\nactual: %v\nexpected: %v\ncidr: %v", actual, tc.expect, tc.input)
		}
	}
}
