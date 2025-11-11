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

	n, _ := strconv.Atoi(line[1:])
	args := make([]string, 0, n)

	for i := 0; i < n; i++ {
		_, _ = r.ReadString('\n')    // skip $len
		val, _ := r.ReadString('\n') // read value
		args = append(args, strings.TrimSpace(val))
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
