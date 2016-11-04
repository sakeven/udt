/*****************************************************************************
Copyright (c) 2001 - 2011, The Board of Trustees of the University of Illinois.
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

* Redistributions of source code must retain the above
  copyright notice, this list of conditions and the
  following disclaimer.

* Redistributions in binary form must reproduce the
  above copyright notice, this list of conditions
  and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the University of Illinois
  nor the names of its contributors may be used to
  endorse or promote products derived from this
  software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*****************************************************************************/

/*****************************************************************************
written by
   Yunhong Gu, last updated 01/18/2011
*****************************************************************************/

#ifndef __UDT_H__
#define __UDT_H__


#ifndef WIN32
   #include <sys/types.h>
   #include <sys/socket.h>
   #include <netinet/in.h>
#else
   #ifdef __MINGW__
      #include <stdint.h>
      #include <ws2tcpip.h>
   #endif
   #include <windows.h>
#endif
#include <fstream>
#include <set>
#include <string>
#include <vector>

#include "udt_base.h"

typedef std::set<UDTSOCKET> ud_set;
#define UD_CLR(u, uset) ((uset)->erase(u))
#define UD_ISSET(u, uset) ((uset)->find(u) != (uset)->end())
#define UD_SET(u, uset) ((uset)->insert(u))
#define UD_ZERO(uset) ((uset)->clear())

class UDT_API CUDTException
{
public:
   CUDTException(int major = 0, int minor = 0, int err = -1);
   CUDTException(const CUDTException& e);
   virtual ~CUDTException();

      // Functionality:
      //    Get the description of the exception.
      // Parameters:
      //    None.
      // Returned value:
      //    Text message for the exception description.

   virtual const char* getErrorMessage();

      // Functionality:
      //    Get the system errno for the exception.
      // Parameters:
      //    None.
      // Returned value:
      //    errno.

   virtual int getErrorCode() const;

      // Functionality:
      //    Clear the error code.
      // Parameters:
      //    None.
      // Returned value:
      //    None.

   virtual void clear();

private:
   int m_iMajor;        // major exception categories

// 0: correct condition
// 1: network setup exception
// 2: network connection broken
// 3: memory exception
// 4: file exception
// 5: method not supported
// 6+: undefined error

   int m_iMinor;		// for specific error reasons
   int m_iErrno;		// errno returned by the system if there is any
   std::string m_strMsg;	// text error message

   std::string m_strAPI;	// the name of UDT function that returns the error
   std::string m_strDebug;	// debug information, set to the original place that causes the error

public: // Error Code
   static const int SUCCESS;
   static const int ECONNSETUP;
   static const int ENOSERVER;
   static const int ECONNREJ;
   static const int ESOCKFAIL;
   static const int ESECFAIL;
   static const int ECONNFAIL;
   static const int ECONNLOST;
   static const int ENOCONN;
   static const int ERESOURCE;
   static const int ETHREAD;
   static const int ENOBUF;
   static const int EFILE;
   static const int EINVRDOFF;
   static const int ERDPERM;
   static const int EINVWROFF;
   static const int EWRPERM;
   static const int EINVOP;
   static const int EBOUNDSOCK;
   static const int ECONNSOCK;
   static const int EINVPARAM;
   static const int EINVSOCK;
   static const int EUNBOUNDSOCK;
   static const int ENOLISTEN;
   static const int ERDVNOSERV;
   static const int ERDVUNBOUND;
   static const int ESTREAMILL;
   static const int EDGRAMILL;
   static const int EDUPLISTEN;
   static const int ELARGEMSG;
   static const int EINVPOLLID;
   static const int EASYNCFAIL;
   static const int EASYNCSND;
   static const int EASYNCRCV;
   static const int ETIMEOUT;
   static const int EPEERERR;
   static const int EUNKNOWN;
};

////////////////////////////////////////////////////////////////////////////////

// If you need to export these APIs to be used by a different language,
// declare extern "C" for them, and add a "udt_" prefix to each API.
// The following APIs: sendfile(), recvfile(), epoll_wait(), geterrormsg(),
// include C++ specific feature, please use the corresponding sendfile2(), etc.

namespace UDT
{

typedef CUDTException ERRORINFO;
typedef UDTOpt SOCKOPT;
typedef CPerfMon TRACEINFO;
typedef ud_set UDSET;

UDT_API extern const UDTSOCKET INVALID_SOCK;
#undef ERROR
UDT_API extern const int ERROR;

UDT_API int startup();
UDT_API int cleanup();
UDT_API UDTSOCKET socket(int af, int type, int protocol);
UDT_API int bind(UDTSOCKET u, const struct sockaddr* name, int namelen);
UDT_API int bind2(UDTSOCKET u, UDPSOCKET udpsock);
UDT_API int listen(UDTSOCKET u, int backlog);
UDT_API UDTSOCKET accept(UDTSOCKET u, struct sockaddr* addr, int* addrlen);
UDT_API int connect(UDTSOCKET u, const struct sockaddr* name, int namelen);
UDT_API int close(UDTSOCKET u);
UDT_API int getpeername(UDTSOCKET u, struct sockaddr* name, int* namelen);
UDT_API int getsockname(UDTSOCKET u, struct sockaddr* name, int* namelen);
UDT_API int getsockopt(UDTSOCKET u, int level, SOCKOPT optname, void* optval, int* optlen);
UDT_API int setsockopt(UDTSOCKET u, int level, SOCKOPT optname, const void* optval, int optlen);
UDT_API int send(UDTSOCKET u, const char* buf, int len, int flags);
UDT_API int recv(UDTSOCKET u, char* buf, int len, int flags);
UDT_API int sendmsg(UDTSOCKET u, const char* buf, int len, int ttl = -1, bool inorder = false);
UDT_API int recvmsg(UDTSOCKET u, char* buf, int len);
UDT_API int64_t sendfile(UDTSOCKET u, std::fstream& ifs, int64_t& offset, int64_t size, int block = 364000);
UDT_API int64_t recvfile(UDTSOCKET u, std::fstream& ofs, int64_t& offset, int64_t size, int block = 7280000);
UDT_API int64_t sendfile2(UDTSOCKET u, const char* path, int64_t* offset, int64_t size, int block = 364000);
UDT_API int64_t recvfile2(UDTSOCKET u, const char* path, int64_t* offset, int64_t size, int block = 7280000);

// select and selectEX are DEPRECATED; please use epoll. 
UDT_API int select(int nfds, UDSET* readfds, UDSET* writefds, UDSET* exceptfds, const struct timeval* timeout);
UDT_API int selectEx(const std::vector<UDTSOCKET>& fds, std::vector<UDTSOCKET>* readfds,
                     std::vector<UDTSOCKET>* writefds, std::vector<UDTSOCKET>* exceptfds, int64_t msTimeOut);

UDT_API int epoll_create();
UDT_API int epoll_add_usock(int eid, UDTSOCKET u, const int* events = NULL);
UDT_API int epoll_add_ssock(int eid, SYSSOCKET s, const int* events = NULL);
UDT_API int epoll_remove_usock(int eid, UDTSOCKET u);
UDT_API int epoll_remove_ssock(int eid, SYSSOCKET s);
UDT_API int epoll_wait(int eid, std::set<UDTSOCKET>* readfds, std::set<UDTSOCKET>* writefds, int64_t msTimeOut,
                       std::set<SYSSOCKET>* lrfds = NULL, std::set<SYSSOCKET>* wrfds = NULL);
UDT_API int epoll_wait2(int eid, UDTSOCKET* readfds, int* rnum, UDTSOCKET* writefds, int* wnum, int64_t msTimeOut,
                        SYSSOCKET* lrfds = NULL, int* lrnum = NULL, SYSSOCKET* lwfds = NULL, int* lwnum = NULL);
UDT_API int epoll_release(int eid);
UDT_API ERRORINFO& getlasterror();
UDT_API int getlasterror_code();
UDT_API const char* getlasterror_desc();
UDT_API int perfmon(UDTSOCKET u, TRACEINFO* perf, bool clear = true);
UDT_API UDTSTATUS getsockstate(UDTSOCKET u);

}  // namespace UDT

#endif
