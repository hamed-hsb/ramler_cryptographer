package bitwise

func And(fByte byte, sByte byte) byte {
	res := fByte & sByte
	return res
}

func Nand(fByte byte, sByte byte) byte {
	res := And(fByte, sByte)
	return ^res
}
func Or(fByte byte, sByte byte) byte {
	res := fByte | sByte
	return res
}
func Nor(fByte byte, sByte byte) byte {
	res := Or(fByte, sByte)
	return ^res
}
func Xor(fByte byte, sByte byte) byte {
	res := fByte ^ sByte
	return res
}
func Xnor(fByte byte, sByte byte) byte {
	res := Xor(fByte, sByte)
	return ^res
}
func Not(byte byte) byte {
	return ^byte
}
