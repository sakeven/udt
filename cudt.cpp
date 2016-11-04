#include <cstdlib>
#include <cstring>
#include <netdb.h>
#include <signal.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <udt.h>
#include "cudt.h"
#include <iostream>

using namespace std;


extern "C" {

  UDT_API const UDTSOCKET INVALID_SOCK=-1;
  UDT_API const int ERROR=-1;

  uint16_t _htons(uint16_t hostshort) {
      return htons(hostshort);
  }

  int udt_startup()
  {
     return UDT::startup();
  }

  int udt_cleanup()
  {
      return UDT::cleanup();
  }

  UDTSOCKET udt_socket(int af, int type, int protocol)
  {
      return UDT::socket(af, type, protocol);
  }

  int udt_bind(UDTSOCKET u, const struct sockaddr* name, int namelen)
  {
      // cout<<namelen<<endl;

      // sockaddr_in addr;
      // addr.sin_family = AF_INET;
      // addr.sin_port = htons(atoi("9000"));
      // addr.sin_addr.s_addr = INADDR_ANY;
      // memset(&(addr.sin_zero), '\0', 8);
      return UDT::bind(u, name, namelen);
      // return UDT::bind(u, (sockaddr*)&addr, sizeof(addr));
  }

  int udt_bind2(UDTSOCKET u, UDPSOCKET udpsock)
  {
      return UDT::bind2(u, udpsock);
  }

  int udt_listen(UDTSOCKET u, int backlog)
  {
      return UDT::listen(u, backlog);
  }

  UDTSOCKET udt_accept(UDTSOCKET u, struct sockaddr* addr, int* addrlen)
  {
      return UDT::accept(u, addr, addrlen);
  }

  int udt_connect(UDTSOCKET u, const struct sockaddr* name, int namelen)
  {
      return UDT::connect(u, name, namelen);
  }

  int udt_close(UDTSOCKET u)
  {
      return UDT::close(u);
  }

  int udt_getpeername(UDTSOCKET u, struct sockaddr* name, int* namelen)
  {
      return UDT::getpeername(u, name, namelen);
  }

  int udt_getsockname(UDTSOCKET u, struct sockaddr* name, int* namelen)
  {
      return UDT::getsockname(u, name, namelen);
  }

  int udt_getsockopt(UDTSOCKET u, int level, SOCKOPT optname, void* optval, int* optlen)
  {
      return UDT::getsockopt(u, level, optname, optval, optlen);
  }

  int udt_setsockopt(UDTSOCKET u, int level, SOCKOPT optname, const void* optval, int optlen)
  {
      return UDT::setsockopt(u, level, optname, optval, optlen);
  }

  int udt_send(UDTSOCKET u, const char* buf, int len, int flags)
  {
      return UDT::send(u, buf, len, flags);
  }

  int udt_recv(UDTSOCKET u, char* buf, int len, int flags)
  {
      return UDT::recv(u, buf, len, flags);
  }

  int udt_sendmsg(UDTSOCKET u, const char* buf, int len, int ttl, int inorder)
  {
      return UDT::sendmsg(u, buf, len, ttl, inorder == 1);
  }

  int udt_recvmsg(UDTSOCKET u, char* buf, int len)
  {
      return UDT::recvmsg(u, buf, len);
  }

  int64_t udt_sendfile2(UDTSOCKET u, const char* path, int64_t* offset, int64_t size, int block)
  {
      return UDT::sendfile2(u, path, offset, size, block);
  }

  int64_t udt_recvfile2(UDTSOCKET u, const char* path, int64_t* offset, int64_t size, int block)
  {
      return UDT::recvfile2(u, path, offset, size, block);
  }

  int udt_epoll_create()
  {
      return UDT::epoll_create();
  }

  int udt_epoll_add_usock(int eid, UDTSOCKET u, const int* events)
  {
      return UDT::epoll_add_usock(eid, u, events);
  }

  int udt_epoll_add_ssock(int eid, SYSSOCKET s, const int* events)
  {
      return UDT::epoll_add_ssock(eid, s, events);
  }

  int udt_epoll_remove_usock(int eid, UDTSOCKET u)
  {
      return UDT::epoll_remove_usock(eid, u);
  }

  int udt_epoll_remove_ssock(int eid, SYSSOCKET s)
  {
      return UDT::epoll_remove_ssock(eid, s);
  }

  int udt_epoll_wait2(int eid, UDTSOCKET* readfds, int* rnum, UDTSOCKET* writefds, int* wnum, int64_t msTimeOut,
                          SYSSOCKET* lrfds, int* lrnum, SYSSOCKET* lwfds, int* lwnum)
  {
      return UDT::epoll_wait2(eid, readfds, rnum, writefds, wnum, msTimeOut, lrfds, lrnum, lwfds, lwnum);
  }

  int udt_epoll_release(int eid)
  {
      return UDT::epoll_release(eid);
  }

  int udt_getlasterror_code()
  {
      return UDT::getlasterror_code();
  }

  const char* udt_getlasterror_desc()
  {
      return UDT::getlasterror_desc();
  }

  int udt_perfmon(UDTSOCKET u, TRACEINFO* perf, int clear)
  {
      return UDT::perfmon(u, perf, clear == 1);
  }

  UDTSTATUS udt_getsockstate(UDTSOCKET u)
  {
      return UDT::getsockstate(u);
  }
}
