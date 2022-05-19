package rawhttp

import (
	"bufio"
	"io"
	"strings"
	"testing"
)

// This testcase panics if it goes wrong
func TestNewRawHeader(t *testing.T) {
	rh := NewRawHeader()
	if len(rh.Names) != 0 {
		t.Fatalf("expected 0 length; got %d", len(rh.Names))
	}
}

func TestRawHeader_ReadHeader_WithBody(t *testing.T) {
	rh := NewRawHeader()
	rb := bufio.NewReader(strings.NewReader(`GET /echoback HTTP/1.1
Host: example.com

`))
	if err := rh.ReadHeader(rb); err != nil {
		t.Fatalf("unexpected read error: %v", err)
	}
	if len(rh.Lines) != 2 {
		t.Errorf("expected %d lines, got %d", 2, len(rh.Lines))
	}
}

func TestRawHeader_ReadHeader_NoBody(t *testing.T) {
	rh := NewRawHeader()
	rb := bufio.NewReader(strings.NewReader(`GET /echoback HTTP/1.1
Host: example.com
`))
	if err := rh.ReadHeader(rb); err != io.EOF {
		t.Fatalf("expected EOF; got %v", err)
	}
	if len(rh.Lines) != 2 {
		t.Errorf("expected %d lines, got %d", 2, len(rh.Lines))
	}
}
