package remoteserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	sockutils "mysock5/sock_utils"
	"net"
)

const (
	CmdConnect = uint8(0x01)
	ATYPIp4    = uint8(0x01)
	ATYPDomain = uint8(0x03)
	ResConnOk  = uint8(0x00)
	ResConnBad = uint8(0x01)
)

type ClientRequestInfo struct {
	version     uint8
	command     uint8
	addressType uint8
	desAddress  []byte
	desPort     []byte
}

func ParseClientRequest(conn *sockutils.ConfusedSocket) (*ClientRequestInfo, error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	buffer := make([]byte, 256)
	requestInfo := &ClientRequestInfo{}
	if length, _ := conn.ReadFull(buffer[:4]); length != 4 {
		return nil, errors.New("get request error")
	}
	requestInfo.version = buffer[0]
	requestInfo.command = buffer[1]
	requestInfo.addressType = buffer[3]
	if requestInfo.addressType == ATYPIp4 {
		if length, _ := conn.ReadFull(buffer[:net.IPv4len]); length != net.IPv4len {
			return nil, errors.New("get ip error")
		}
		requestInfo.desAddress = append(requestInfo.desAddress, buffer[:net.IPv4len]...)
	} else if requestInfo.addressType == ATYPDomain {
		conn.ReadFull(buffer[:1])
		domainLength := buffer[0]
		if length, _ := conn.ReadFull(buffer[:domainLength]); length != int(domainLength) {
			return nil, errors.New("get domain error")
		}
		requestInfo.desAddress = append(requestInfo.desAddress, buffer[:domainLength]...)
	} else {
		return nil, errors.New("invalid addr type")
	}
	if length, _ := conn.ReadFull(buffer[:2]); length != 2 {
		return nil, errors.New("get port error")
	}
	requestInfo.desPort = append(requestInfo.desPort, buffer[:2]...)
	return requestInfo, nil
}

func AttachConnect(conn *sockutils.ConfusedSocket) (*net.TCPConn, error) {

	requestInfo, err := ParseClientRequest(conn)
	fmt.Println(*requestInfo)
	if err != nil {
		return nil, err
	}
	var addr string
	if requestInfo.addressType == ATYPIp4 {
		addr = fmt.Sprintf("%d.%d.%d.%d", requestInfo.desAddress[0],
			requestInfo.desAddress[1], requestInfo.desAddress[2], requestInfo.desAddress[3])
	} else if requestInfo.addressType == ATYPDomain {
		addr = string(requestInfo.desAddress)
	}
	port := binary.BigEndian.Uint16(requestInfo.desPort)
	destAddr := fmt.Sprintf("%s:%d", addr, port)
	destTcpAddr, err1 := net.ResolveTCPAddr("tcp", destAddr)
	if err1 != nil {
		err = conn.Write([]byte{SocksV5, ResConnBad, 0x00, ATYPIp4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		return nil, err1
	}
	destConn, connError := net.DialTCP("tcp", nil, destTcpAddr)
	if connError != nil {
		err = conn.Write([]byte{SocksV5, ResConnBad, 0x00, ATYPIp4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		return nil, connError
	}
	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// connectInfo := destConn.LocalAddr()
	err = conn.Write([]byte{SocksV5, ResConnOk, 0x00, ATYPIp4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		return nil, errors.New("write response error")
	}
	return destConn, nil

}

func Forwardata(src *sockutils.ConfusedSocket, des *net.TCPConn) {
	exitChan := make(chan bool, 1)
	go func(exit chan bool) {
		buf := make([]byte, 4096)
		defer src.Close()
		defer des.Close()
		for {
			length, err := src.Read(buf)
			if err != nil {
				exit <- true
				return
			}
			if _, err := des.Write(buf[:length]); err != nil {
				exit <- true

				return
			}
		}
	}(exitChan)
	go func(exit chan bool) {
		buf := make([]byte, 4096)
		defer src.Close()
		defer des.Close()
		for {
			length, err := des.Read(buf)
			if err != nil {
				exit <- true
				return
			}
			if err := src.Write(buf[:length]); err != nil {
				exit <- true
				return
			}
		}
	}(exitChan)
	<-exitChan
}
