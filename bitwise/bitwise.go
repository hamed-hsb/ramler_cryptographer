/*
Name: bitwise Package
Auth: Nadali Khalili
Date: 14 Dec 2021
This pakage is prepared for bitwise operation. We've implemented XOR, OR, And, XNOR, NAND, NOR for two bytes.
*/
package bitwise

/**
@param fByte: A byte type argument which you need as the first operand of AND operator.
@param sByte: A byte type argument which you need as the second operand of AND operator.
@return: A byte type value which specify the result of the operation

This functioin get two bytes and do AND operation on them, then return the result.
*/
func And(fByte byte, sByte byte) byte {
	res := fByte & sByte
	return res
}

/**
@param fByte: A byte type argument which you need as the first operand of Nand operator.
@param sByte: A byte type argument which you need as the second operand of Nand operator.
@return: A byte type value which specify the result of the operation

This functioin get two bytes and do Nand operation on them, then return the result.
*/
func Nand(fByte byte, sByte byte) byte {
	res := And(fByte, sByte)
	return ^res
}

/**
@param fByte: A byte type argument which you need as the first operand of Or operator.
@param sByte: A byte type argument which you need as the second operand of Or operator.
@return: A byte type value which specify the result of the operation

This functioin get two bytes and do Or operation on them, then return the result.
*/
func Or(fByte byte, sByte byte) byte {
	res := fByte | sByte
	return res
}

/**
@param fByte: A byte type argument which you need as the first operand of Xor operator.
@param sByte: A byte type argument which you need as the second operand of Xor operator.
@return: A byte type value which specify the result of the operation

This functioin get two bytes and do Xor operation on them, then return the result.
*/
func Nor(fByte byte, sByte byte) byte {
	res := Or(fByte, sByte)
	return ^res
}

/**
@param fByte: A byte type argument which you need as the first operand of Xor operator.
@param sByte: A byte type argument which you need as the second operand of Xor operator.
@return: A byte type value which specify the result of the operation

This functioin get two bytes and do Xor operation on them, then return the result.
*/
func Xor(fByte byte, sByte byte) byte {
	res := fByte ^ sByte
	return res
}

/**
@param fByte: A byte type argument which you need as the first operand of Xnor operator.
@param sByte: A byte type argument which you need as the second operand of Xnor operator.
@return: A byte type value which specify the result of the operation

This functioin get two bytes and do Xnor operation on them, then return the result.
*/
func Xnor(fByte byte, sByte byte) byte {
	res := Xor(fByte, sByte)
	return ^res
}

/**
@param byte: A byte type argument which you need as the first operand of Xnor operator.
@return: A byte type value which specify the result of the operation

This functioin get a byte and do Not operation on it, then return the result.
*/
func Not(byte byte) byte {
	return ^byte
}
