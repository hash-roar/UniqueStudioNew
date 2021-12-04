package remoteserver

import (
	"errors"
	"fmt"
	sockutils "mysock5/sock_utils"
)

var RegisterUserPass map[string]string

func init() {
	RegisterUserPass = make(map[string]string)
	RegisterUserPass["user1"] = "pass1"
}

const (
	SocksV5             = uint8(0x05)
	AuthNone            = uint8(0x00)
	AuthUserPass        = uint8(0x02)
	AuthNoUsable        = uint8(0xFF)
	AuthResultOk        = uint8(0x00)
	AuthResultNoUser    = uint8(0x01)
	AuthResultWrongPass = uint8(0x02)
)

func ParseHeader(conn *sockutils.ConfusedSocket) (uint8, error) {
	// +----+----------+----------+
	// |VER | NUMMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	header := make([]byte, 255)
	length, _ := conn.ReadFull(header[:2])
	fmt.Println("header:  ", header[:2])
	sockVersion := header[0]
	if length != 2 {
		return 0, errors.New("parse header error")
	}
	methodNum := header[1]
	length, _ = conn.ReadFull(header[:methodNum])
	if length != int(methodNum) {
		return 0, errors.New("parse methods error")
	}
	return sockVersion, nil
}

func AuthClient(conn *sockutils.ConfusedSocket) (bool, error) {
	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	if err := conn.Write([]byte{SocksV5, AuthUserPass}); err != nil {
		return false, err
	}
	version, name, pass, err := getUserInfo(conn)
	if err != nil {
		return false, err
	}
	if passValue, ok := RegisterUserPass[name]; ok {
		if pass == passValue {
			conn.Write([]byte{version, AuthResultOk})
			return true, nil
		} else {
			conn.Write([]byte{version, AuthResultWrongPass})
		}
	} else {
		conn.Write([]byte{version, AuthResultNoUser})
	}
	return false, nil

}

func getUserInfo(conn *sockutils.ConfusedSocket) (version byte, name string, pass string, err error) {

	// VERSION 	USERNAME_LENGTH 	USERNAME 	PASSWORD_LENGTH 	PASSWORD
	// 1字节 	    1字节 		         1-255字节 	 1字节 				 1-255字节
	// 0x01 	    0x01 	             0x0a 	    0x01 	            0x0a
	buffer := make([]byte, 256)
	if length, _ := conn.ReadFull(buffer[:1]); length != 1 {
		err = errors.New("get version error")
		return
	}
	version = buffer[0]
	if length, _ := conn.ReadFull(buffer[:1]); length != 1 {
		err = errors.New("get user name length error")
		return
	}
	nameLength := buffer[0]
	if length, _ := conn.ReadFull(buffer[:nameLength]); length != int(nameLength) {
		err = errors.New("get name error")
		return
	}
	name = string(buffer[0:nameLength])

	if length, _ := conn.ReadFull(buffer[:1]); length != 1 {
		err = errors.New("get user pass length error")
		return
	}
	passLength := buffer[0]
	if length, _ := conn.ReadFull(buffer[:passLength]); length != int(passLength) {
		err = errors.New("get user pass error")
		return
	}
	pass = string(buffer[:passLength])
	err = nil
	return
}
