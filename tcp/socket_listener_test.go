package tcp

import (
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestListen(t *testing.T) {
	ssl := &SocketListener{
		Address: "localhost",
		Port:    "9000",
	}

	ssl.Listen()
}

func TestWrite(t *testing.T) {
	// path := "192.168.64.18:32074"
	path := "127.0.0.1:9000"

	// Connect to socket server
	conn, err := net.DialTimeout("tcp", path, 5*time.Second)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	i := 0
	for {
		// Send data
		data := "99 hello socket, this is message " + strconv.Itoa(i) + " from 1"
		i++
		_, _ = conn.Write([]byte(data))
		time.Sleep(1000 * time.Microsecond)
	}
}

func TestWrite2(t *testing.T) {
	path := "192.168.64.18:32074"
	path = "127.0.0.1:9000"

	// Connect to socket server
	conn, err := net.DialTimeout("tcp", path, 5*time.Second)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	i := 0
	for {
		// Send data
		data := "99 hello socket, this is message " + strconv.Itoa(i) + " from 2"
		i++
		n, err := conn.Write([]byte(data))
		logrus.Info(n, err)
		time.Sleep(10000 * time.Microsecond)
	}
}
