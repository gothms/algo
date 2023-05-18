package bit

// ReverseBits 翻转位
// Go 源码：bits.Reverse
func ReverseBits(x uint) uint {
	x = x>>16 | x<<16
	x = (x&0xff00ff00)>>8 | (x&0x00ff00ff)<<8
	x = (x&0xf0f0f0f0)>>4 | (x&0x0f0f0f0f)<<4 // f：1111 0：0000
	x = (x&0xcccccccc)>>2 | (x&0x33333333)<<2 // c：1100 3：0011
	x = (x&0xaaaaaaaa)>>1 | (x&0x55555555)<<1 // a：1010 5：0101
	return x
}
