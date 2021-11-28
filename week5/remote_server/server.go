package remoteserver

import (
	"errors"
	"fmt"
	"io"
	"log"
	sockUtils "mysock5/sock_utils"
	"net"
	"os"
)

type RemoteServerConf struct {
	LisenAddr *net.TCPAddr
}

type RemoteServer struct {
	Listener        *net.TCPListener
	LocalServerConn *net.TCPConn //can be designed as array
	Conf            RemoteServerConf
}

type ConnPipe struct {
	LocalServerConn *net.TCPConn
	DestConn        *net.TCPConn
}

func main() {
	service := "127.0.0.1:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	sockUtils.CheckErrorFatal(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	sockUtils.CheckErrorFatal(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			sockUtils.CheckErrorLog(err)
			continue
		}
		go ConnHandle(conn)
	}
}

func NewRemoteServer(conf *RemoteServerConf) (*RemoteServer, error) {
	if conf.LisenAddr == nil {
		return nil, errors.New("conf error")
	}
	return &RemoteServer{Conf: *conf}, nil
}

func (s *RemoteServer) ListenLocalServer() {
	listener, err := net.ListenTCP("tcp", s.Conf.LisenAddr)
	if err != nil {
		log.Fatal("init remote server error")
	}
	s.Listener = listener
}

func (s *RemoteServer) AcceptLocalConn() {
	defer s.Listener.Close()
	for {
		localServerConn, err := s.Listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		go ConnHandle(localServerConn)
	}
}

func ConnHandle(conn *net.TCPConn) {
	defer conn.Close()
	if version, err := ParseHeader(conn); err != nil || version != SocksV5 {
		log.Println(err)
		return
	}
	if ok, err := AuthClient(conn); !ok && err != nil {
		log.Println(err)
		return
	}
	destConn, err := AttachConnect(conn)
	if err != nil {
		log.Println(err)
		return
	}
	Forwardata(conn, destConn)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func CopyData(localConn net.Conn, destConn net.Conn) {
	go func() {
		defer localConn.Close()
		defer destConn.Close()
		io.Copy(destConn, localConn)
	}()
	go func() {
		defer localConn.Close()
		defer destConn.Close()
		io.Copy(localConn, destConn)
	}()
}
