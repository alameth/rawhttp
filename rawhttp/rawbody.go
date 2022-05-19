package rawhttp

import "bufio"

type RawBody []string

func (this *RawBody) ReadBody(br *bufio.Reader) error {
	for {
		line, err := readline(br)
		if err != nil {
			return err
		}
		*this = append(*this, line)
	}
}
