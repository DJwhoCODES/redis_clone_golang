package proto

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Parse(r *bufio.Reader) ([]string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSpace(line)

	if !strings.HasPrefix(line, "*") {
		return strings.Split(line, " "), nil
	}

	n, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid array length: %v", err)
	}

	args := make([]string, 0, n)

	for i := 0; i < n; i++ {
		lenLine, err := r.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("failed to read bulk length: %v", err)
		}
		lenLine = strings.TrimSpace(lenLine)

		if !strings.HasPrefix(lenLine, "$") {
			return nil, fmt.Errorf("expected bulk string, got: %s", lenLine)
		}

		length, err := strconv.Atoi(lenLine[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid bulk string length: %v", err)
		}

		buf := make([]byte, length+2)
		if _, err := r.Read(buf); err != nil {
			return nil, fmt.Errorf("failed to read bulk data: %v", err)
		}

		args = append(args, string(buf[:length]))
	}

	return args, nil
}

func WriteSimple(conn net.Conn, msg string) {
	conn.Write([]byte("+" + msg + "\r\n"))
}

func WriteError(conn net.Conn, msg string) {
	conn.Write([]byte("-ERR " + msg + "\r\n"))
}

func WriteBulk(conn net.Conn, msg string) {
	conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(msg), msg)))
}

func WriteInt(conn net.Conn, n int) {
	conn.Write([]byte(fmt.Sprintf(":%d\r\n", n)))
}

func WriteNull(conn net.Conn) {
	conn.Write([]byte("$-1\r\n"))
}
