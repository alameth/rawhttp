package rawhttp

import (
	"bufio"
	"fmt"
	"io"
)

// The RawHeader object holds an HTTP Header as a slice of strings, indexed
// (but not parsed) by header field name. Its purpose is to support read,
// minimal modifications. and faithful write of an HTTP header no matter
// how badly malformed.

type RawHeader struct {
	Lines []string
	Names NameMap
}

type NameMap map[string][]string

// NewRawHeader returns a reference to a newly instantiated RawHeader object.
func NewRawHeader() RawHeader {
	return RawHeader{Names: make(NameMap)}
}

// ReadHeader reads an HTTP header from the specified reader stream. If a blank
// line is read, it is discarded and the method returns with error set to nil.
// If an EOF is reached, the method returns with error set to io.EOF.
func (this *RawHeader) ReadHeader(scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			return nil
		}
		this.Lines = append(this.Lines, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	} else {
		return io.EOF
	}
}

// SetContentLength finds the first occurrence of the Content-Length field
// in the header, and replaces it with a new field with the specified value.
// If the header has no Content-Length fields, one is appended to the end.
func (this *RawHeader) SetContentLength(length int) {
	newField := fmt.Sprintf("Content-Length: %d", length)
	if _, ok := this.Names["content-length"]; ok {

	} else {
		this.Lines = append(this.Lines, newField)
	}
}

// SetHost
func (this *RawHeader) SetHost(servername string) {
	newField := fmt.Sprintf("Host: %s", servername)
	if _, ok := this.Names["content-length"]; ok {

	} else {
		this.Lines = append(this.Lines, newField)
	}
}

// WriteHeader streams the entire header to the specified writer. The caller
// is responsible for writing the blank line between the header and body.
func (this RawHeader) WriteHeader(w io.Writer) error {
	for _, line := range this.Lines {
		if n, err := fmt.Fprintf(w, "%s\r\n", line); err != nil {
			return err
		} else if n != len(line)+2 {
			return fmt.Errorf("length error: tried %d, wrote %d", len(line), n)
		}
	}
	w.Write([]byte("\r\n"))
	return nil
}

func (this *RawHeader) addField(name, value string) {

}
