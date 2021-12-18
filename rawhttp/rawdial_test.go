package rawhttp

import "testing"

func TestBuildConnectName(t *testing.T) {
	tests := []struct {
		inhost  string
		intls   bool
		exphost string
		expport string
		exptls  bool
	}{
		{"example.com", false, "example.com", "80", false},
		{"example.com:80", false, "example.com", "80", false},
		{"example.com:443", false, "example.com", "443", false},
		{"http://example.com", false, "example.com", "80", false},
		{"https://example.com", false, "example.com", "443", true},
		{"https://example.com/", false, "example.com", "443", true},
		{"http://bilge.com:27/", true, "bilge.com", "27", true},
	}
	for _, tc := range tests {
		gothost, gotport, gottls, err := getHostPort(tc.inhost, tc.intls)
		if err != nil {
			t.Fatalf("[%s] unexpected error: %v", tc.inhost, err)
		}
		if gothost != tc.exphost {
			t.Errorf("[%s] got %s, expected %s", tc.inhost, gothost, tc.exphost)
		}
		if gotport != tc.expport {
			t.Errorf("[%s] got %s, expected %s", tc.inhost, gotport, tc.expport)
		}
		if gottls != tc.exptls {
			t.Errorf("[%s] got TLS %v, expected TLS %v", tc.inhost, gottls, tc.exptls)
		}
	}
}

func TestBuildConnectName_Negative(t *testing.T) {
	tests := []struct {
		inhost  string
		intls   bool
	}{
		{"example.com:", false},
		{"https://example.com:", false},
	}
	for _, tc := range tests {
		gothost, gotport, _, err := getHostPort(tc.inhost, tc.intls)
		if err == nil {
			t.Errorf("[%s] expected error; got %s %s", tc.inhost, gothost, gotport)
		}
	}
}
