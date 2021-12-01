package sockutils

import (
	"io"
	"net"
)

var Rc4SecKey string = "thisfuckingjesus"

type ConfusedSocket struct {
	CipherIn  *Rc4Cipher //用于向加密通道写
	Cipherout *Rc4Cipher
	conn      *net.TCPConn
}

func NewConfusionSock(conn *net.TCPConn) (*ConfusedSocket, error) {
	return &ConfusedSocket{CipherIn: NewCipher(), Cipherout: NewCipher(), conn: conn}, nil
}

func (c *ConfusedSocket) ReadFull(buf []byte) (length int, err error) {

	if length, err = io.ReadFull(c.conn, buf); length != len(buf) {
		return
	}
	c.Cipherout.DecodeBuf(buf)
	return
}
func (c *ConfusedSocket) Read(buf []byte) (length int, err error) {

	if length, err = c.conn.Read(buf); err != nil {
		return length, err
	}
	c.Cipherout.DecodeBuf(buf[0:length])
	return length, nil
}

func (c *ConfusedSocket) Write(buf []byte) error {
	c.CipherIn.EncodeBuf(buf)
	if _, err := c.conn.Write(buf); err != nil {
		return err
	}
	return nil
}

func (c *ConfusedSocket) Close() {
	c.conn.Close()
}

func (c *ConfusedSocket) GetConn() *net.TCPConn {
	return c.conn
}
func (c *ConfusedSocket) GetCipherIn() *Rc4Cipher {
	return c.CipherIn
}
func (c *ConfusedSocket) GetCipherout() *Rc4Cipher {
	return c.Cipherout
}
