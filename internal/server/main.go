package server

import (
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
	defer listener.Close()

	err = server.listen(listener)
	if err != nil {
		return nil, err
	}
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
			res, err := s.handle(c)
			if err != nil {
				log.Printf("[ERROR] - %v", res)
			}
			log.Printf("%s", res)
		}(conn)
	}
}

func (s *Server) handle(conn net.Conn) (Response, error) {
	response := "HTTP/1.1 200 OK\r\n" +
		"Content-Type: text/plain\r\n" +
		"Content-Length: 13\r\n" +
		"\r\n" +
		"Hello World!"
	return Response(response), nil
}
