package server

import (
	"net"
	"strings"

	"github.com/djwhocodes/redis-clone/internal/proto"
	"github.com/djwhocodes/redis-clone/internal/store"
)

type Handler struct {
	store *store.Store
}

func NewHandler(s *store.Store) *Handler {
	return &Handler{store: s}
}

func (h *Handler) Handle(conn net.Conn, args []string) {
	if len(args) == 0 {
		proto.WriteError(conn, "empty command")
		return
	}

	cmd := strings.ToUpper(args[0])

	switch cmd {
	case "PING":
		proto.WriteSimple(conn, "PONG")

	case "SET":
		if len(args) != 3 {
			proto.WriteError(conn, "wrong number of arguments for SET")
			return
		}
		h.store.Set(args[1], args[2])
		proto.WriteSimple(conn, "OK")

	case "GET":
		if len(args) != 2 {
			proto.WriteError(conn, "wrong number of arguments for GET")
			return
		}
		val, ok := h.store.Get(args[1])
		if !ok {
			proto.WriteBulk(conn, "")
			return
		}
		proto.WriteBulk(conn, val)

	case "DEL":
		if len(args) != 2 {
			proto.WriteError(conn, "wrong number of arguments for DEL")
			return
		}
		h.store.Del(args[1])
		proto.WriteSimple(conn, "OK")

	default:
		proto.WriteError(conn, "unknown command "+cmd)
	}
}
