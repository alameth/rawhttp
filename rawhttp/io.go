package rawhttp

import (
	"bufio"
	"fmt"
)

func readline(br *bufio.Reader) (string, error) {
	line, err := br.ReadString('\n')
	if err != nil {
		if len(line) != 0 {
			// Last line not newline terminated
			return "", fmt.Errorf("unexpected EOF")
		}
		return "", err
	}
	return line, nil
}
