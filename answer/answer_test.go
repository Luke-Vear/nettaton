package main

import (
	"testing"
)

var jrm = []byte(`{
		"answer": "10.126.72.254",
		"ipAddress": "10.126.72.171",
		"network": "255.255.255.0",
		"questionKind": "last"
	}`)

func TestFirst(t *testing.T) {
	type args struct {
		ip  string
		net string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "first-prefix-30",
			args: args{ip: "10.72.243.104", net: "/30"},
			want: "10.72.243.101",
		},
		{
			name: "first-prefix-24",
			args: args{ip: "10.217.75.42", net: "/24"},
			want: "10.217.75.1",
		},
		{
			name: "first-prefix-22",
			args: args{ip: "10.90.149.74", net: "/22"},
			want: "10.90.148.1",
		},
		{
			name: "first-prefix-16",
			args: args{ip: "10.170.101.236", net: "/16"},
			want: "10.170.0.1",
		},
		{
			name: "first-prefix-12",
			args: args{ip: "10.28.244.157", net: "/12"},
			want: "10.16.0.1",
		},
		{
			name: "first-netmask-255.255.255.224",
			args: args{ip: "10.218.18.110", net: "255.255.255.224"},
			want: "10.218.18.97",
		},
		{
			name: "first-netmask-255.255.255.0",
			args: args{ip: "10.200.103.119", net: "255.255.255.0"},
			want: "10.200.103.1",
		},
		{
			name: "first-netmask-255.255.192.0",
			args: args{ip: "10.4.198.63", net: "255.255.192.0"},
			want: "10.4.192.1",
		},
		{
			name: "first-netmask-255.255.0.0",
			args: args{ip: "10.33.60.245", net: "255.255.0.0"},
			want: "10.33.0.1",
		},
		{
			name: "first-netmask-255.248.0.0",
			args: args{ip: "10.250.253.191", net: "255.248.0.0"},
			want: "10.248.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := First(tt.args.ip, tt.args.net); got != tt.want {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args struct {
		ip  string
		net string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	/* TODO: want's calculating manually
	{
		name: "last-prefix-30",
		args: args{ip: "10.72.243.104", net: "/30"},
		want: "10.72.243.103",
	},
	{
		name: "last-prefix-24",
		args: args{ip: "10.217.75.42", net: "/24"},
		want: "10.217.75.254",
	},
	{
		name: "last-prefix-22",
		args: args{ip: "10.90.149.74", net: "/22"},
		want: "10.90.151.254",
	},
	{
		name: "last-prefix-16",
		args: args{ip: "10.170.101.236", net: "/16"},
		want: "10.170.0.1",
	},
	{
		name: "last-prefix-12",
		args: args{ip: "10.28.244.157", net: "/12"},
		want: "10.16.0.1",
	},
	{
		name: "last-netmask-255.255.255.224",
		args: args{ip: "10.218.18.110", net: "255.255.255.224"},
		want: "10.218.18.97",
	},
	{
		name: "last-netmask-255.255.255.0",
		args: args{ip: "10.200.103.119", net: "255.255.255.0"},
		want: "10.200.103.1",
	},
	{
		name: "last-netmask-255.255.192.0",
		args: args{ip: "10.4.198.63", net: "255.255.192.0"},
		want: "10.4.192.1",
	},
	{
		name: "last-netmask-255.255.0.0",
		args: args{ip: "10.33.60.245", net: "255.255.0.0"},
		want: "10.33.0.1",
	},
	{
		name: "last-netmask-255.248.0.0",
		args: args{ip: "10.250.253.191", net: "255.248.0.0"},
		want: "10.248.0.1",
	},
	*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.ip, tt.args.net); got != tt.want {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBroadcast(t *testing.T) {
	type args struct {
		ip  string
		net string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Broadcast(tt.args.ip, tt.args.net); got != tt.want {
				t.Errorf("Broadcast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		ip  string
		net string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Range(tt.args.ip, tt.args.net); got != tt.want {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
