package paddingworker

import "bytes"

//PCKS5 padding
func PCKS5(src []byte) ([]byte, int) {
	return PCKS(src, 8)
}

//PCKS7  padding
func PCKS7(src []byte) ([]byte, int) {
	return PCKS(src, 256)
}

//PCKS custom padding with block size
func PCKS(src []byte, blocksize int) ([]byte, int) {
	padlen := (blocksize - len(src)%blocksize)

	if padlen < blocksize {
		pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
		rt := append(src, pad...)
		return rt, (len(rt))
	}
	return src, len(src)

}

// RemovePadding remove padding
func RemovePadding(src []byte, blocksize int) []byte {
	if (len(src) % blocksize) == 0 {
		return src
	}
	padlen := int(src[len(src)-1])
	return src[0 : len(src)-padlen]

}
