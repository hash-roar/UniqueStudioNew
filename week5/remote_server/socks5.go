package remoteserver

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
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

func ParseClientRequest(conn *net.TCPConn) (*ClientRequestInfo, error) {
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	buffer := make([]byte, 256)
	requestInfo := &ClientRequestInfo{}
	if length, _ := io.ReadFull(conn, buffer[:4]); length != 4 {
		return nil, errors.New("get request error")
	}
	requestInfo.version = buffer[0]
	requestInfo.command = buffer[1]
	requestInfo.addressType = buffer[3]
	if requestInfo.addressType == ATYPIp4 {
		if length, _ := io.ReadFull(conn, buffer[:net.IPv4len]); length != net.IPv4len {
			return nil, errors.New("get ip error")
		}
		requestInfo.desAddress = append(requestInfo.desAddress, buffer[:net.IPv4len]...)
	} else if requestInfo.addressType == ATYPDomain {
		io.ReadFull(conn, buffer[:1])
		domainLength := buffer[0]
		if length, _ := io.ReadFull(conn, buffer[:domainLength]); length != int(domainLength) {
			return nil, errors.New("get domain error")
		}
		requestInfo.desAddress = append(requestInfo.desAddress, buffer[:domainLength]...)
	} else {
		return nil, errors.New("invalid addr type")
	}
	if length, _ := io.ReadFull(conn, buffer[:2]); length != 2 {
		return nil, errors.New("get port error")
	}
	requestInfo.desPort = append(requestInfo.desPort, buffer[:2]...)
	return requestInfo, nil
}

func AttachConnect(conn *net.TCPConn) (*net.TCPConn, error) {

	requestInfo, err := ParseClientRequest(conn)
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
		_, err = conn.Write([]byte{SocksV5, ResConnBad, 0x00, ATYPIp4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		return nil, err1
	}
	destConn, connError := net.DialTCP("tcp", nil, destTcpAddr)
	if connError != nil {
		_, err = conn.Write([]byte{SocksV5, ResConnBad, 0x00, ATYPIp4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
		return nil, connError
	}
	// +----+-----+-------+------+----------+----------+
	// |VER | REP |  RSV  | ATYP | BND.ADDR | BND.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// connectInfo := destConn.LocalAddr()
	_, err = conn.Write([]byte{SocksV5, ResConnOk, 0x00, ATYPIp4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	if err != nil {
		return nil, errors.New("write response error")
	}
	return destConn, nil

}

func Forwardata(src *net.TCPConn, des *net.TCPConn) {
	go io.Copy(src, des)
	go io.Copy(des, src)
}
