UDT.go A Golang wrapper for the UDT project
=================================================

UDP-based Data Transfer Protocol (UDT), is a high-performance data transfer protocol
designed for transferring large volumetric datasets over high-speed wide area networks.
Such settings are typically disadvantageous for the more common TCP protocol.
The original C++ UDT project can be found on SourceForge http://udt.sourceforge.net/

### Getting Started

First You'll need to compile the UDT C source code.  To do this run the build script
that is located in the root directory of the project.  You'll need g++ on your
system to successfully compile the UDT C library.

load dy
```shell
# osx
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:$PWD/udt4/src
# linux
export LD_LIBRARY_PATH=$PWD/udt4/src:$LD_LIBRARY_PATH
```
