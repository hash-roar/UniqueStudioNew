package sockutils

var S [256]uint8
var keyLen uint8
var keyBytes []uint8

type Rc4Cipher struct {
	i uint8
	j uint8
	S [256]uint8
}

func swap(a *uint8, b *uint8) {
	temp := *a
	*a = *b
	*b = temp
}

func init() {
	keyLen = uint8(len(Rc4SecKey))
	keyBytes = []byte(Rc4SecKey)
}

func NewCipher() *Rc4Cipher {
	r := &Rc4Cipher{}
	r.i, r.j = 0, 0
	for i := uint(0); i < 256; i++ {
		r.S[i] = uint8(i)
	}
	j := uint8(0)
	for i := uint(0); i < 256; i++ {
		r.j = (r.j + r.S[r.i])
		j = (j + r.S[i] + keyBytes[i%uint(keyLen)])
		swap(&r.S[i], &r.S[j])
	}
	return r
}

func (r *Rc4Cipher) Encode(src byte) byte {
	r.i = r.i + 1
	r.j = r.j + r.S[r.i]
	swap(&r.S[r.i], &r.S[r.j])
	des := src ^ r.S[r.S[r.i]+r.S[r.j]]
	return des
}

func (r *Rc4Cipher) Decode(src uint8) uint8 {
	r.i = (r.i + 1)
	r.j = (r.j + r.S[r.i])
	swap(&r.S[r.i], &r.S[r.j])
	des := src ^ r.S[r.S[r.i]+r.S[r.j]]
	return des
}

func (r *Rc4Cipher) EncodeBuf(buf []byte) {
	for i, _ := range buf {
		r.i = (r.i + 1)
		r.j = (r.j + r.S[r.i])
		swap(&r.S[r.i], &r.S[r.j])
		buf[i] = buf[i] ^ r.S[r.S[r.i]+r.S[r.j]]
	}
}
func (r *Rc4Cipher) DecodeBuf(buf []byte) {
	for i, _ := range buf {
		r.i = (r.i + 1)
		r.j = (r.j + r.S[r.i])
		swap(&r.S[r.i], &r.S[r.j])
		buf[i] = buf[i] ^ r.S[r.S[r.i]+r.S[r.j]]
	}
}
