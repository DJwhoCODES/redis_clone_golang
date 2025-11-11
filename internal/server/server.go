package server

import (
	"bufio"
	"net"

	"github.com/djwhocodes/redis-clone/internal/logger"
	"github.com/djwhocodes/redis-clone/internal/proto"
)

type Server struct {
	addr    string
	handler *Handler
}

func NewServer(addr string, handler *Handler) *Server {
	return &Server{addr: addr, handler: handler}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	logger.Info.Println("Server started on", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error.Println("Accept error:", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	logger.Info.Println("New connection from", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		args, err := proto.Parse(reader)
		if err != nil {
			logger.Info.Println("Client disconnected:", conn.RemoteAddr())
			return
		}
		s.handler.Handle(conn, args)
	}
}
