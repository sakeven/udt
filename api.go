package udt

// #cgo CXXFLAGS: -I${SRCDIR}/udt4/src
// #cgo LDFLAGS: -L${SRCDIR}/udt4/src -ludt -lstdc++ -lpthread -lm
// #include <stdlib.h>
// #include <string.h>
// #include <arpa/inet.h>
// #include "cudt.h"
import "C"

import (
	"errors"
	"unsafe"
	"fmt"
)

func Startup() error {
	n, err := C.udt_startup()
	if err != nil {
		return err
	}

	if n != 0 {
		return GetLastError()
	}

	return nil
}

func Cleanup() error {
	n, err := C.udt_cleanup()
	if err != nil {
		return err
	}

	if n != 0 {
		return GetLastError()
	}

	return nil
}

const (
	AF_INET  = int(C.AF_INET)
	AF_INET6 = int(C.AF_INET6)

	SOCK_STREAM = int(C.SOCK_STREAM)
	SOCK_DGRAM  = int(C.SOCK_DGRAM)
)

func Socket(af int, _type int) (C.UDTSOCKET, error) {
	s, err := C.udt_socket(C.int(af), C.int(_type), C.int(0))
	if err != nil {
		return C.INVALID_SOCK, err
	}

	if s == C.INVALID_SOCK {
		return s, GetLastError()
	}

	return s, nil
}

func Bind(u C.UDTSOCKET, name *C.struct_sockaddr_in) error {
	n := C.udt_bind(u, (*C.struct_sockaddr)(unsafe.Pointer(name)), C.int(unsafe.Sizeof(*name)))
	if n != 0 {
		return GetLastError()
	}

	return nil
}

func Bind2(u C.UDTSOCKET, udpsock C.UDPSOCKET) error {
	n := C.udt_bind2(u, udpsock)
	if n != 0 {
		return GetLastError()
	}

	return nil
}

func Listen(u C.UDTSOCKET, backlog int) error {
	n, err := C.udt_listen(u, C.int(backlog))
	if err != nil {
		return err
	}

	if n != 0 {
		return GetLastError()
	}

	return nil
}

func Accept(u C.UDTSOCKET, addr *C.struct_sockaddr_in, addrlen *int) (C.UDTSOCKET, error) {
	sock := C.udt_accept(u, (*C.struct_sockaddr)(unsafe.Pointer(addr)), (*C.int)(unsafe.Pointer(addrlen)))
	if sock == C.INVALID_SOCK {
		return sock, GetLastError()
	}

	return sock, nil
}

func Connect(u C.UDTSOCKET, name *C.struct_sockaddr_in) error {
	n := C.udt_connect(u, (*C.struct_sockaddr)(unsafe.Pointer(name)), C.int(unsafe.Sizeof(*name)))
	if n != 0 {
		return GetLastError()
	}

	return nil
}

func Close(u C.UDTSOCKET) error {
	n := C.udt_close(u)
	if n != 0 {
		return GetLastError()
	}

	return nil
}

func GetPeername(u C.UDTSOCKET, name *C.struct_sockaddr, namelen *int) int {
	return int(C.udt_getpeername(u, name, (*C.int)(unsafe.Pointer(namelen))))
}

func GetSockname(u C.UDTSOCKET, name *C.struct_sockaddr, namelen *int) int {
	return int(C.udt_getsockname(u, name, (*C.int)(unsafe.Pointer(namelen))))
}

func GetSockopt(u C.UDTSOCKET, level int, optname C.SOCKOPT, optval unsafe.Pointer, optlen *int) int {
	return int(C.udt_getsockopt(u, C.int(level), optname, optval, (*C.int)(unsafe.Pointer(optlen))))
}

func SetSockopt(u C.UDTSOCKET, level int, optname C.SOCKOPT, optval unsafe.Pointer, optlen int) int {
	return int(C.udt_setsockopt(u, C.int(level), optname, optval, C.int(optlen)))
}

func Send(u C.UDTSOCKET, buf []byte) (int, error) {
	n, err := C.udt_send(u, (*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)), 0)
	if err != nil {
		return 0, err
	}

	if n == C.ERROR {
		return 0, GetLastError()
	}

	return int(n), nil
}

func Recv(u C.UDTSOCKET, buf []byte) (int, error) {
	n, err := C.udt_recv(u, (*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)), 0)
	if err != nil {
		return 0, err
	}

	if n == C.ERROR {
		return 0, GetLastError()
	}

	return int(n), nil
}

func SendMsg(u C.UDTSOCKET, buf []byte, ttl int, inorder int) int {
	return int(C.udt_sendmsg(u, (*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)), C.int(ttl), C.int(inorder)))
}

func RecvMsg(u C.UDTSOCKET, buf []byte) int {
	return int(C.udt_recvmsg(u, (*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf))))
}

func Sendfile(u C.UDTSOCKET, path []byte, offset *int64, size int64, block int) error {
	n := C.udt_sendfile2(u, (*C.char)(unsafe.Pointer(&path[0])), (*C.int64_t)(unsafe.Pointer(offset)), C.int64_t(size), C.int(block))
	if n == C.int64_t(C.ERROR) {
		return GetLastError()
	}

	return nil
}

func Recvfile(u C.UDTSOCKET, path []byte, offset *int64, size int64, block int) error {
	n := C.udt_recvfile2(u, (*C.char)(unsafe.Pointer(&path[0])), (*C.int64_t)(unsafe.Pointer(offset)), C.int64_t(size), C.int(block))
	if n == C.int64_t(C.ERROR) {
		return GetLastError()
	}

	return nil
}

// int udt_epoll_create();
// int udt_epoll_add_usock(int eid, u C.UDTSOCKET, const int* events);
// int udt_epoll_add_ssock(int eid, SYSSOCKET s, const int* events);
// int udt_epoll_remove_usock(int eid, u C.UDTSOCKET);
// int udt_epoll_remove_ssock(int eid, SYSSOCKET s);
// int udt_epoll_wait2(int eid, UDTSOCKET* readfds, int* rnum, UDTSOCKET* writefds, int* wnum, int64_t msTimeOut,
//                         SYSSOCKET* lrfds, int* lrnum, SYSSOCKET* lwfds, int* lwnum);
// int udt_epoll_release(int eid);
func Perfmon(u C.UDTSOCKET, perf *C.TRACEINFO, clear int) error {
	n := C.udt_perfmon(u, perf, C.int(clear))
	if n == C.ERROR {
		return GetLastError()
	}

	return nil
}

func GetLastErrorCode() int {
	return int(C.udt_getlasterror_code())
}

func GetLastError() error {
	return errors.New(GetLastErrorDesc())
}

func GetLastErrorDesc() string {
	code := GetLastErrorCode()
	desc := C.udt_getlasterror_desc()
	return fmt.Sprintf("code %d: %s", code, C.GoString(desc))
}

func GetSockState(u C.UDTSOCKET) C.enum_UDTSTATUS {
	return C.udt_getsockstate(u)
}
