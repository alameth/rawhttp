package rawhttp

import (
	"bufio"
	"fmt"
	"io"
)

type RawHTTP struct {
	Header RawHeader
	Body   RawBody
}

func NewRawHTTP() RawHTTP {
	return RawHTTP{
		Header: NewRawHeader(),
		Body:   []string{},
	}
}

func (this *RawHTTP) Read(r io.Reader) error {
	br := bufio.NewReader(r)
	if err := this.Header.ReadHeader(br); err == io.EOF {
		return nil
	} else if err != nil {
		return err
	}
	if err := this.Body.ReadBody(br); err != nil && err != io.EOF {
		return err
	}
	return nil
}

func (this *RawHTTP) Write(w io.Writer) error {
	bw := bufio.NewWriter(w)
	if err := this.Header.WriteHeader(bw); err != nil {
		return err
	}
	for _, line := range this.Body {
		if n, err := bw.WriteString(line); err != nil {
			return err
		} else if n != len(line) {
			return fmt.Errorf("length error: tried %d, wrote %d", len(line), n)
		}
	}
	bw.Flush()
	return nil
}
