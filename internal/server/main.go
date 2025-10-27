package server

import (
	"GO_HTTP_PROTOCOL/internal/headers"
	"GO_HTTP_PROTOCOL/internal/request"
	"GO_HTTP_PROTOCOL/internal/response"
	"bytes"
	"log"
	"net"
	"strconv"
)

type Server struct {
	listener net.Listener
	handler  Handler
}

func Serve(port int, handler Handler) (*Server, error) {

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	server := Server{
		listener: listener,
		handler:  handler,
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
			s.handle(c)
		}(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	req, err := request.RequestFromReader(conn)

	if err != nil {
		hErr := &HandleError{
			StatusCode: response.HTTP_400,
			Message:    err.Error(),
		}
		hErr.Write(conn)
		return
	}

	defaultResponse := &response.Response{
		Status:  response.HTTP_200,
		Body:    *bytes.NewBuffer([]byte{}),
		Headers: headers.NewHeaders(),
	}

	hErr := s.handler(defaultResponse, req)
	if hErr != nil {
		hErr := &HandleError{
			StatusCode: response.HTTP_400,
			Message:    err.Error(),
		}
		hErr.Write(conn)
		return
	}
	err = response.WriteStatusLine(conn, defaultResponse.Status)
	if err != nil {
		hErr := &HandleError{
			StatusCode: response.HTTP_500,
			Message:    err.Error(),
		}
		hErr.Write(conn)
		return
	}
	headers := defaultResponse.GetHeaders()
	err = response.WriteHeaders(conn, headers)
	if err != nil {
		hErr := &HandleError{
			StatusCode: response.HTTP_500,
			Message:    err.Error(),
		}
		hErr.Write(conn)
		return
	}
	defaultResponse.Write(conn)

}
