package rawhttp

import (
	"bufio"
	"fmt"
	"io"
)

type RawHTTP struct {
	Header RawHeader
	Body   []string
}

func NewRawHTTP() RawHTTP {
	return RawHTTP{
		Header: NewRawHeader(),
		Body:   []string{},
	}
}

func (this *RawHTTP) Read(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	if err := this.Header.ReadHeader(scanner); err == io.EOF {
		return nil
	} else if err != nil {
		return err
	}
	for scanner.Scan() {
		this.Body = append(this.Body, scanner.Text())
	}
	return scanner.Err()
}

func (this *RawHTTP) Write(w io.Writer) error {
	bw := bufio.NewWriter(w)
	if err := this.Header.WriteHeader(bw); err != nil {
		return err
	}
	for _, line := range this.Body {
		if n, err := fmt.Fprintf(bw, "%s\r\n", line); err != nil {
			return err
		} else if n != len(line)+2 {
			return fmt.Errorf("length error: tried %d, wrote %d", len(line), n)
		}
	}
	bw.Flush()
	return nil
}
