package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestParseNetwork(t *testing.T) {
	tests := map[string]struct {
		address    string
		network    string
		wantErr    bool
		errMessage string
	}{
		"valid_class_c": {
			address: "192.168.1.1",
			network: "/24",
			wantErr: false,
		},
		"invalid_ip": {
			address:    "300.168.1.1",
			network:    "/24",
			wantErr:    true,
			errMessage: "invalid IP address",
		},
		"invalid_network": {
			address:    "192.168.1.1",
			network:    "/33",
			wantErr:    true,
			errMessage: "invalid network mask",
		},
		"malformed_network": {
			address:    "192.168.1.1",
			network:    "abc",
			wantErr:    true,
			errMessage: "invalid network mask",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := parseNetwork(tc.address, tc.network)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tc.errMessage)
				}
				if !strings.Contains(err.Error(), tc.errMessage) {
					t.Fatalf("expected error containing %q, got %q", tc.errMessage, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

func TestPrintNetworkAddress(t *testing.T) {
	tests := map[string]struct {
		address string
		network string
		wantOut string
		wantErr bool
	}{
		"class_c": {
			address: "192.168.1.1",
			network: "/24",
			wantOut: `Network address: 192.168.1.0
Usable host IP range: 192.168.1.1 - 192.168.1.254
Broadcast address: 192.168.1.255
Total number of hosts: 256
Number of useable hosts: 254
Subnet mask: 255.255.255.0
Wildcard mask: 0.0.0.255
`,
		},
		"tiny_network": {
			address: "10.0.0.1",
			network: "/30",
			wantOut: `Network address: 10.0.0.0
Usable host IP range: 10.0.0.1 - 10.0.0.2
Broadcast address: 10.0.0.3
Total number of hosts: 4
Number of useable hosts: 2
Subnet mask: 255.255.255.252
Wildcard mask: 0.0.0.3
`,
		},
		"class_a": {
			address: "10.0.0.1",
			network: "/8",
			wantOut: `Network address: 10.0.0.0
Usable host IP range: 10.0.0.1 - 10.255.255.254
Broadcast address: 10.255.255.255
Total number of hosts: 16777216
Number of useable hosts: 16777214
Subnet mask: 255.0.0.0
Wildcard mask: 0.255.255.255
`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Capture stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			err := printNetworkAddress(tc.address, tc.network)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read captured output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			got := buf.String()

			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.wantOut {
				t.Fatalf("got:\n%s\nwant:\n%s", got, tc.wantOut)
			}
		})
	}
}

func TestIPv4MaskString(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  string
	}{
		"valid_24": {
			input: []byte{255, 255, 255, 0},
			want:  "255.255.255.0",
		},
		"valid_16": {
			input: []byte{255, 255, 0, 0},
			want:  "255.255.0.0",
		},
		"invalid_length": {
			input: []byte{255, 255},
			want:  "invalid mask",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := ipv4MaskString(tc.input)
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}

func TestWildcard(t *testing.T) {
	tests := map[string]struct {
		input []byte
		want  string
	}{
		"mask_24": {
			input: []byte{255, 255, 255, 0},
			want:  "0.0.0.255",
		},
		"mask_16": {
			input: []byte{255, 255, 0, 0},
			want:  "0.0.255.255",
		},
		"mask_8": {
			input: []byte{255, 0, 0, 0},
			want:  "0.255.255.255",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := wildcard(tc.input)
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
