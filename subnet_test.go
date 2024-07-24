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
		"104.28.48.74": {address: "104.28.48.74", network: "/30", want: "104.28.48.72"}}

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

func TestUseableHostIPRange(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		want    string
	}{
		"248.192.215.107": {address: "248.192.215.107", network: "/30", want: "248.192.215.105 - 248.192.215.106"},
		"38.73.20.159":    {address: "38.73.20.159", network: "/23", want: "38.73.20.1 - 38.73.21.254"},
		"179.241.4.46":    {address: "179.241.4.46", network: "17", want: "179.241.0.1 - 179.241.127.254"},
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
