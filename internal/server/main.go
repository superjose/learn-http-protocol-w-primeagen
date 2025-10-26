package server

import (
	"GO_HTTP_PROTOCOL/internal/response"
	"log"
	"net"
	"strconv"
)

type Server struct {
	listener net.Listener
}

type Response string

func Serve(port int) (*Server, error) {

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	server := Server{
		listener: listener,
	}

	if err != nil {
		return nil, err
	}

	go func() {
		err = server.listen(listener)
		if err != nil {
			log.Printf("[ERROR] - %v", err)
		}
	}()
	return &server, nil
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) listen(listener net.Listener) error {

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			err := s.handle(c)
			if err != nil {
				log.Printf("[ERROR] - %v", err)
			}

		}(conn)
	}
}

func (s *Server) handle(conn net.Conn) error {
	defer conn.Close()
	err := response.WriteStatusLine(conn, response.HTTP_200)
	if err != nil {
		return err
	}
	headers := response.GetDefaultHeaders(0)
	err = response.WriteHeaders(conn, headers)
	if err != nil {
		return err
	}
	return nil
}
