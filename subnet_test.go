package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNetworkAddress(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    string
	}{
		"class_c_network": {
			address: "192.168.1.1",
			network: "/24",
			want:    "192.168.1.0",
		},
		"small_subnet": {
			address: "104.28.48.74",
			network: "/30",
			want:    "104.28.48.72",
		},
		"large_subnet": {
			address: "10.0.0.1",
			network: "/8",
			want:    "10.0.0.0",
		},
		"medium_subnet": {
			address: "172.16.1.1",
			network: "/16",
			want:    "172.16.0.0",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := networkAddress(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestUsableHostIPRange(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    string
	}{
		"tiny_network": {
			address: "248.192.215.107",
			network: "/30",
			want:    "248.192.215.105 - 248.192.215.106",
		},
		"small_network": {
			address: "38.73.20.159",
			network: "/23",
			want:    "38.73.20.1 - 38.73.21.254",
		},
		"medium_network": {
			address: "179.241.4.46",
			network: "/17",
			want:    "179.241.0.1 - 179.241.127.254",
		},
		"large_network": {
			address: "10.0.0.1",
			network: "/8",
			want:    "10.0.0.1 - 10.255.255.254",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := usableHostIPRange(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestBroadcastAddress(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    string
	}{
		"class_c": {
			address: "192.168.1.1",
			network: "/24",
			want:    "192.168.1.255",
		},
		"small_subnet": {
			address: "179.241.4.46",
			network: "/30",
			want:    "179.241.4.47",
		},
		"medium_subnet": {
			address: "179.241.4.46",
			network: "/17",
			want:    "179.241.127.255",
		},
		"large_subnet": {
			address: "10.0.0.1",
			network: "/8",
			want:    "10.255.255.255",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := broadcastAddress(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestTotalNumberOfHosts(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    uint32
	}{
		"tiny_network": {
			address: "179.241.4.46",
			network: "/30",
			want:    4,
		},
		"small_network": {
			address: "192.168.1.1",
			network: "/24",
			want:    256,
		},
		"medium_network": {
			address: "172.16.1.1",
			network: "/16",
			want:    65536,
		},
		"large_network": {
			address: "10.0.0.1",
			network: "/8",
			want:    16777216,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := totalNumberOfHosts(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestNumberOfUsableHosts(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    uint32
	}{
		"tiny_network": {
			address: "179.241.4.46",
			network: "/30",
			want:    2,
		},
		"small_network": {
			address: "192.168.1.1",
			network: "/24",
			want:    254,
		},
		"medium_network": {
			address: "172.16.1.1",
			network: "/16",
			want:    65534,
		},
		"large_network": {
			address: "10.0.0.1",
			network: "/8",
			want:    16777214,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := numberOfUsableHosts(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestSubnetMask(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    string
	}{
		"30_bit_mask": {
			address: "179.241.4.46",
			network: "/30",
			want:    "255.255.255.252",
		},
		"24_bit_mask": {
			address: "192.168.1.1",
			network: "/24",
			want:    "255.255.255.0",
		},
		"16_bit_mask": {
			address: "172.16.1.1",
			network: "/16",
			want:    "255.255.0.0",
		},
		"8_bit_mask": {
			address: "10.0.0.1",
			network: "/8",
			want:    "255.0.0.0",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := subnetMask(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestWildcardMask(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    string
	}{
		"30_bit_mask": {
			address: "179.241.4.46",
			network: "/30",
			want:    "0.0.0.3",
		},
		"24_bit_mask": {
			address: "192.168.1.1",
			network: "/24",
			want:    "0.0.0.255",
		},
		"16_bit_mask": {
			address: "172.16.1.1",
			network: "/16",
			want:    "0.0.255.255",
		},
		"8_bit_mask": {
			address: "10.0.0.1",
			network: "/8",
			want:    "0.255.255.255",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := wildcardMask(tc.address, tc.network)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
