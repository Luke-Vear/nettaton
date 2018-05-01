package quiz

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update golden record files?")

type netw struct {
	IP      string
	Network string
}

type output struct {
	Kind     string
	Solution string
}

type outputs []output

func (o outputs) Len() int           { return len(o) }
func (o outputs) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o outputs) Less(i, j int) bool { return o[i].Kind < o[j].Kind }

type record struct {
	Input   netw
	Outputs outputs
}

func TestQuizGolden(t *testing.T) {

	tests := []struct {
		index string
		input netw
	}{
		{"001", netw{"10.10.10.10", "255.255.255.252"}},
		{"002", netw{"10.10.10.10", "255.255.255.0"}},
		{"003", netw{"10.10.10.10", "255.240.0.0"}},

		{"004", netw{"172.16.222.222", "255.255.255.252"}},
		{"005", netw{"172.16.222.222", "255.255.255.0"}},
		{"006", netw{"172.16.222.222", "255.240.0.0"}},

		{"007", netw{"192.168.123.123", "255.255.255.252"}},
		{"008", netw{"192.168.123.123", "255.255.255.0"}},
		{"009", netw{"192.168.123.123", "255.240.0.0"}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s", test.index), func(t *testing.T) {

			var outs outputs

			for kind := range solvers {
				solution := NewQuestion(test.input.IP, test.input.Network, kind).Solution()
				outs = append(outs, output{
					kind,
					solution,
				})
			}
			sort.Sort(outs)

			actual := record{
				Input:   test.input,
				Outputs: outs,
			}

			goldenFileName := fmt.Sprintf("testdata/QuizGolden.%s.json", test.index)

			if *update {
				bs, err := json.MarshalIndent(actual, "", "  ")
				if err != nil {
					t.Error(err)
				}

				err = ioutil.WriteFile(goldenFileName, bs, 0666)
				if err != nil {
					t.Error(err)
				}

				return
			}

			goldenBs, err := ioutil.ReadFile(goldenFileName)
			if err != nil {
				t.Error(err)
			}

			var expected record

			err = json.Unmarshal(goldenBs, &expected)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, expected, actual)
		})
	}
}

func TestNewAnswerer(t *testing.T) {
	a := newAnswerer(&Question{
		ID:      "abc",
		IP:      "10.0.0.0",
		Network: "24",
		Kind:    "first",
	})

	assert.NotNil(t, a)
}

func TestToCidr(t *testing.T) {
	netmask := "255.240.0.0"
	netmaskOut := toCidr(netmask)

	assert.Equal(t, "12", netmaskOut)

	cidr := "12"
	cidrOut := toCidr(cidr)

	assert.Equal(t, "12", cidrOut)
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

func TestValidQuestionKind(t *testing.T) {
	assert.True(t, ValidQuestionKind("first"))
	assert.False(t, ValidQuestionKind("smarch"))
}
