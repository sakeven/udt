package udt

import (
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

func TestInitializationAndShutdown(t *testing.T) {
	e := Startup()
	if e != nil {
		t.Fatal(e)
	}

	e = Cleanup()
	if e != nil {
		t.Fatal(e)
	}

}

func TestTransferFile(t *testing.T) {
	Startup()
	defer Cleanup()

	s, err := ListenUDT("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	fpath := "test.file"
	stat, err := os.Stat(fpath)
	size := stat.Size()

	go func() {
		defer s.Close()
		f, _ := ioutil.ReadFile(fpath)
		for i := 1; i <= 10; i++ {
			time.Sleep(1 * time.Second)
			now := time.Now()

			sock, err := Dial("ip4", "127.0.0.1:9000")
			if err != nil {
				t.Errorf("Unable to dial: %s", err)
				return
			}

			var offset int64 = 0
			for offset < size {
				r := offset + 7320000
				if r > size {
					r = size
				}
				n, err := sock.Write(f[offset:r])
				if err != nil {
					log.Printf("write error %s", err)
					break
				}

				offset += int64(n)
			}

			log.Printf("Transfer time %d %d", time.Now().Sub(now)/time.Second, sock.sock)
			perf, err := sock.Perfmon(0)
			if err != nil {
				t.Error(err)
				return
			}
			log.Printf("Transfer file speed = %f Mbits/sec", float64(perf.mbpsSendRate))
			sock.Close()
		}
	}()

	for {
		sock, err := s.Accept()
		if err != nil {
			return
		}

		go func(sock *Socket) {
			defer sock.Close()
			log.Printf("Get Socket %d", sock.sock)

			file, err := os.OpenFile("xxx", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				t.Error(err)
				return
			}
			defer file.Close()

			var offset int64 = 0
			msg := make([]byte, 7320000)

			for offset < size {
				n, err := sock.Read(msg)
				if err != nil {
					log.Printf("read error %s", err)
					return
				}

				_, err = file.Write(msg[:n])
				if err != nil {
					log.Printf("write error %s", err)
				}
				offset += int64(n)
			}
		}(sock)
	}
}

func TestSendfileRecvfile(t *testing.T) {
	Startup()
	defer Cleanup()

	s, err := ListenUDT("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	fpath := "test.file"
	stat, err := os.Stat(fpath)
	size := stat.Size()
	go func() {
		defer s.Close()
		for i := 1; i <= 10; i++ {
			now := time.Now()
			sock, err := Dial("ip4", "127.0.0.1:9000")
			if err != nil {
				t.Errorf("Unable to dial: %s", err)
				return
			}

			var offset int64 = 0
			sock.Sendfile(fpath, &offset, size)

			log.Printf("Sendfile time %d %d", time.Now().Sub(now)/time.Second, sock.sock)

			perf, err := sock.Perfmon(0)
			if err != nil {
				t.Error(err)
				return
			}

			log.Printf("Sendfile speed = %f Mbits/sec", float64(perf.mbpsSendRate))

			sock.Close()
			time.Sleep(1 * time.Second)
		}
	}()
	for true {
		sock, err := s.Accept()
		if err != nil {
			return
		}

		go func(sock *Socket) {
			defer sock.Close()
			log.Printf("Get Socket %d", sock.sock)
			
			var offset int64 = 0
			sock.Recvfile("xxx", &offset, size)
		}(sock)
	}
}

func TestReadWrite(t *testing.T) {
	Startup()
	defer Cleanup()

	s, err := ListenUDT("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	go func() {
		buf := []byte(testMsg)
		for i := 1; i <= 10; i++ {
			sock, err := Dial("ip4", "127.0.0.1:9000")
			if err != nil {
				t.Errorf("Unable to dial: %s", err)
				return
			}

			sock.Write(buf)
			sock.Close()
			time.Sleep(time.Second)
		}
	}()
	buf := make([]byte, len(testMsg))
	for {
		sock, err := s.Accept()
		if err != nil {
			return
		}

		go func(sock *Socket) {
			defer sock.Close()
			sock.Read(buf)
		}(sock)
	}
}

const testMsg = `
func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}
func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}

func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}

func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}func TestDail(t *testing.T) {
	s, err := Dial("ip4", "127.0.0.1:9000")
	if err != nil {
		t.Errorf("Unable to dial: %s", err)
		return
	}

	log.Printf("Socket is: %s", s)
}

`
