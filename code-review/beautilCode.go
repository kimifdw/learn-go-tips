package main

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxCons  int
	TLS      *tls.Config
}

type Option func(*Server)

func Protocol(proto string) Option {
	return func(s *Server) {
		s.Protocol = proto
	}
}

func Timeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.Timeout = timeout
	}
}

func MaxConns(maxconns int) Option {
	return func(server *Server) {
		server.MaxCons = maxconns
	}
}

func TLS(tls *tls.Config) Option {
	return func(server *Server) {
		server.TLS = tls
	}
}

func NewServer(addr string, port int) (*Server, error) {
	return nil, nil
}

func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
	return nil, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
	return nil, nil
}

func NewTLSServerWithMaxAndTimeout(addr string, port int, maxconns int) (*Server, error) {
	return nil, nil
}
