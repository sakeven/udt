package udt

// #include <stdlib.h>
// #include <string.h>
// #include <arpa/inet.h>
// #include "cudt.h"
import "C"

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unsafe"
)

type Socket struct {
	sock C.UDTSOCKET
}

func Dial(network string, address string) (socket *Socket, err error) {
	var af int

	switch network {
	case "ip4":
		af = AF_INET
	case "ip6":
		af = AF_INET6
	default:
		return nil, fmt.Errorf("network must be either ip4 or ip6")
	}

	sock, err := Sokcet(af, SOCK_STREAM)
	if err != nil {
		return nil, err
	}

	serv_addr, err := addr(af, address)
	if err != nil {
		return nil, err
	}

	if err := Connect(sock, &serv_addr); err != nil {
		return nil, fmt.Errorf("connect error %s", err)
	}

	socket = &Socket{
		sock: sock,
	}

	return
}

func ListenUDT(network string, address string) (socket *Socket, err error) {
	var af int

	switch network {
	case "ip4":
		af = AF_INET
	case "ip6":
		af = AF_INET6
	default:
		return nil, fmt.Errorf("network must be either ip4 or ip6")
	}

	sock, err := Sokcet(af, SOCK_STREAM)
	if err != nil {
		return nil, err
	}

	serv_addr, err := addr(af, address)
	if err != nil {
		return nil, err
	}

	if err := Bind(sock, &serv_addr); err != nil {
		log.Printf("Bind error")
		return nil, err
	}

	if err := Listen(sock, 1024); err != nil {
		log.Printf("Listen error")
		return nil, err
	}

	socket = &Socket{
		sock: sock,
	}
	return
}

func (s *Socket) Accept() (*Socket, error) {
	var clientaddr C.struct_sockaddr_in
	var addrlen int
	sock, err := Accept(s.sock, &clientaddr, &addrlen)
	if err != nil {
		return nil, err
	}

	return &Socket{
		sock: sock,
	}, nil
}

func (s *Socket) Write(buf []byte) (int, error) {
	return Send(s.sock, buf)
}

func (s *Socket) Read(buf []byte) (int, error) {
	return Recv(s.sock, buf)
}

func (s *Socket) Close() error {
	err := Close(s.sock)
	if err != nil {
		log.Printf("close error %s %d %d", err, s.sock, C.INVALID_SOCK)
	}
	return err
}

func (s *Socket) Perfmon(clear int) (C.TRACEINFO, error) {
	var perf C.TRACEINFO
	err := Perfmon(s.sock, &perf, clear)

	return perf, err
}

func (s *Socket) Sendfile(path string, offset *int64, size int64) error {
	return Sendfile(s.sock, []byte(path), offset, size, 7320000)
}

func (s *Socket) Recvfile(path string, offset *int64, size int64) error{
	return Recvfile(s.sock, []byte(path), offset, size, 7320000)
}

func addr(af int, address string) (serv_addr C.struct_sockaddr_in, err error) {

	splitAddr := strings.Split(address, ":")
	if len(splitAddr) != 2 {
		return serv_addr, fmt.Errorf("Please specify an address as host:port")
	}

	host, _port := splitAddr[0], splitAddr[1]
	port, err := strconv.Atoi(_port)
	if err != nil {
		return serv_addr, fmt.Errorf("Invalid port: %s", _port)
	}
	serv_addr.sin_family = C.sa_family_t(af)
	serv_addr.sin_port = C.in_port_t(C._htons(C.uint16_t(port)))
	chost := C.CString(host)
	defer C.free(unsafe.Pointer(chost))

	if _, err := C.inet_pton(C.int(af), chost, unsafe.Pointer(&serv_addr.sin_addr)); err != nil {
		return serv_addr, fmt.Errorf("Unable to convert IP address: %s", err)
	}

	if _, err := C.memset(unsafe.Pointer(&(serv_addr.sin_zero)), 0, 8); err != nil {
		return serv_addr, fmt.Errorf("Unable to zero sin_zero")
	}

	return
}
