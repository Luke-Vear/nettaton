package subnet

/*
import (
	"testing"
)

var jrm = []byte(`{
		"answer": "10.126.72.254",
		"ipAddress": "10.126.72.171",
		"network": "255.255.255.0",
		"questionKind": "last"
	}`)

// type ClientRequest struct {
// 	Answer       string `json:"answer"`
// 	IPAddress    string `json:"ipAddress"`
// 	Network      string `json:"network"`
// 	QuestionKind string `json:"questionKind"`
// }

func TestParseAnswer(t *testing.T) {
	testTable := []struct {
		cr        *ClientRequest
		expected  string
		errExpect error
	}{
		{
			cr: &ClientRequest{
				Answer:       "10.72.243.105",
				IPAddress:    "10.72.243.104",
				Network:      "30",
				QuestionKind: "first",
			},
			expected:  "10.72.243.105",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.217.75.1",
				IPAddress:    "10.217.75.42",
				Network:      "24",
				QuestionKind: "first",
			},
			expected:  "10.217.75.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.90.148.1",
				IPAddress:    "10.90.149.74",
				Network:      "22",
				QuestionKind: "first",
			},
			expected:  "10.90.148.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.170.0.1",
				IPAddress:    "10.170.101.236",
				Network:      "16",
				QuestionKind: "first",
			},
			expected:  "10.170.0.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.16.0.1",
				IPAddress:    "10.28.244.157",
				Network:      "12",
				QuestionKind: "first",
			},
			expected:  "10.16.0.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.218.18.97",
				IPAddress:    "10.218.18.110",
				Network:      "255.255.255.224",
				QuestionKind: "first",
			},
			expected:  "10.218.18.97",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.200.103.1",
				IPAddress:    "10.200.103.119",
				Network:      "255.255.255.0",
				QuestionKind: "first",
			},
			expected:  "10.200.103.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.4.192.1",
				IPAddress:    "10.4.198.63",
				Network:      "255.255.192.0",
				QuestionKind: "first",
			},
			expected:  "10.4.192.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.33.0.1",
				IPAddress:    "10.33.60.245",
				Network:      "255.255.0.0",
				QuestionKind: "first",
			},
			expected:  "10.33.0.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.248.0.1",
				IPAddress:    "10.250.253.191",
				Network:      "255.248.0.0",
				QuestionKind: "first",
			},
			expected:  "10.248.0.1",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.72.243.106",
				IPAddress:    "10.72.243.104",
				Network:      "30",
				QuestionKind: "last",
			},
			expected:  "10.72.243.106",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.217.75.254",
				IPAddress:    "10.217.75.42",
				Network:      "24",
				QuestionKind: "last",
			},
			expected:  "10.217.75.254",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.90.151.254",
				IPAddress:    "10.90.149.74",
				Network:      "22",
				QuestionKind: "last",
			},
			expected:  "10.90.151.254",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.170.255.254",
				IPAddress:    "10.170.101.236",
				Network:      "16",
				QuestionKind: "last",
			},
			expected:  "10.170.255.254",
			errExpect: nil,
		},
		{
			cr: &ClientRequest{
				Answer:       "10.31.255.254",
				IPAddress:    "10.28.244.157",
				Network:      "12",
				QuestionKind: "last",
			},
			expected:  "10.31.255.254",
			errExpect: nil,
		},
	}
	for _, tt := range testTable {
		actual, errActual := parseAnswer(tt.cr)
		if errActual != tt.errExpect || actual != tt.expected {
			t.Errorf("actual: %v\nexpected: %v\ncr: %v\nerrExpected: %v\nerrActual: %v\n",
				actual, tt.expected, tt.cr, tt.errExpect, errActual)
		}
	}
}

// {
// 	name: "last-netmask-255.255.255.224",
// 	args: args{ip: "10.218.18.110", net: "255.255.255.224"},
// 	want: "10.218.18.97",
// },
// {
// 	name: "last-netmask-255.255.255.0",
// 	args: args{ip: "10.200.103.119", net: "255.255.255.0"},
// 	want: "10.200.103.254",
// },
// {
// 	name: "last-netmask-255.255.192.0",
// 	args: args{ip: "10.4.198.63", net: "255.255.192.0"},
// 	want: "10.4.255.254",
// },
// {
// 	name: "last-netmask-255.255.0.0",
// 	args: args{ip: "10.33.60.245", net: "255.255.0.0"},
// 	want: "10.33.255.254",
// },
// {
// 	name: "last-netmask-255.248.0.0",
// 	args: args{ip: "10.250.253.191", net: "255.248.0.0"},
// 	want: "10.255.255.254",
// },

func TestBroadcast(t *testing.T) {
	// {
	// 	name: "broadcast-prefix-30",
	// 	args: args{ip: "10.72.243.104", net: "30"},
	// 	want: "10.72.243.107",
	// },
	// {
	// 	name: "broadcast-prefix-24",
	// 	args: args{ip: "10.217.75.42", net: "24"},
	// 	want: "10.217.75.255",
	// },
	// {
	// 	name: "broadcast-prefix-22",
	// 	args: args{ip: "10.90.149.74", net: "22"},
	// 	want: "10.90.151.255",
	// },
	// {
	// 	name: "broadcast-prefix-16",
	// 	args: args{ip: "10.170.101.236", net: "16"},
	// 	want: "10.170.255.255",
	// },
	// {
	// 	name: "broadcast-prefix-12",
	// 	args: args{ip: "10.28.244.157", net: "12"},
	// 	want: "10.31.255.255",
	// },
	// {
	// 	name: "broadcast-netmask-255.255.255.224",
	// 	args: args{ip: "10.218.18.110", net: "255.255.255.224"},
	// 	want: "10.218.18.98",
	// },
	// {
	// 	name: "broadcast-netmask-255.255.255.0",
	// 	args: args{ip: "10.200.103.119", net: "255.255.255.0"},
	// 	want: "10.200.103.255",
	// },
	// {
	// 	name: "broadcast-netmask-255.255.192.0",
	// 	args: args{ip: "10.4.198.63", net: "255.255.192.0"},
	// 	want: "10.4.255.255",
	// },
	// {
	// 	name: "broadcast-netmask-255.255.0.0",
	// 	args: args{ip: "10.33.60.245", net: "255.255.0.0"},
	// 	want: "10.33.255.255",
	// },
	// {
	// 	name: "broadcast-netmask-255.248.0.0",
	// 	args: args{ip: "10.250.253.191", net: "255.248.0.0"},
	// 	want: "10.255.255.255",
	// },

}

func TestRange(t *testing.T) {

}
*/
