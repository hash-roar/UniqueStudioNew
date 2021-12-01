package sockutils

import (
	"log"
	"testing"
)

func TestCipher(t *testing.T) {
	encodeCipher := NewCipher()
	decodeCipher := NewCipher()
	str := "this is for test string"
	buf := []byte(str)
	encodeCipher.EncodeBuf(buf)
	log.Println(string(buf))
	decodeCipher.DecodeBuf(buf)
	if string(buf) != str {
		t.Errorf("expect %s,but get %s", str, string(buf))
	}
}
