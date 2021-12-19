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
@param byte: A byte type argument which you need as the first operand of Not operator.
@return: A byte type value which specify the result of the operation

This functioin get a byte and do Not operation on it, then return the result.
*/
func Not(byte byte) byte {
	return ^byte
}

/**
@param byte: A byte type argument which you need as the first operand of Not operator.
@return: A byte type value which specify the result of the operation

This functioin get a bit and do Not operation on it, then return the result.
The difference between Not and NotBit is that Not operatoration on 0 is 11111111 and its 255 on decimal but in NotBit it's going to be 1
*/
func NotBit(bit uint8) uint8 {
	if bit == 1 {
		return 0
	} else {
		return 1
	}
}

/**
@param str: A string type argument which you need to make operati.
@return: A boolian and an string type values which specify the status of operation and the result of the operation.

This function get an 8 bits string of and do XOR on each couple of bits and returns a 4 bit binary string as result.
*/
func MakeXorCoupleBits(str string) (string, bool) {
	if len(str) != 8 {
		return "you string should has 8 bits", true
	}
	var hashed []uint8
	for i := 0; i < 8; i += 2 {
		hashed = append(hashed, Xor(str[i]-48, str[i+1]-48)+48)
	}
	return string(hashed), false
}

/**

@param str: A string type argument which you need to make operati.
@return: A boolian and an string type values which specify the status of operation and the result of the operation.

This function get an 8 bits string and do NAND and AND on each couple of bits and returns a 4 bit binary string as result.
*/
func MakeNandAndCoupleBits(str string) (string, bool) {
	if len(str) != 8 {
		return "you string should has 8 bits", true
	}
	nand := true
	var hashed []uint8
	for i := 0; i < 8; i += 2 {
		if nand {
			hashed = append(hashed, NotBit(And(str[i]-48, str[i+1]-48))+48)
			nand = false
		} else {
			hashed = append(hashed, And(str[i]-48, str[i+1]-48)+48)
			nand = true
		}
	}
	return string(hashed), false
}
