/**
This package convert a string to bits and convert bits to array of 8 bits
This package get a string argument and char and then convert to byte and then convert to bits
This package get a string of bits and separate to 8 bits
*/

package bitSeparator

/**
@Au: Mohammad Rezaie
@Date: 13 Dec 2021
*/

import (
	"fmt"
	"reflect"
)

/**
this func get a string and convert it to an array of characters
*/
func strToChar(str string) ([]string, error ) {
	if len(str) <= 0 { //error handling
		err := fmt.Errorf("string is empty!")
		return nil, err
	}

	var charArr []string // for storing array of character
	chars := []rune(str) // for converting string to characters using rune function
	for i := 0; i < len(chars); i++ {
		char := string(chars[i]) // for each character that we want to store it to "charArr"
		charArr = append(charArr, char) // storing characters
	}
	return charArr , nil
}

/**
this func convert characters to an array of byte
*/
func charToByte(chars []string) ([]byte, error) {
	if len(chars) <= 0 { //error handling
		err := fmt.Errorf("srting is empty!")
		return nil, err
	}

	var byteArr []byte // for storing array of bytes
	for i := 0; i < len(chars); i++ {
		bytes := []byte(chars[i]) // for storing bytes of a character
		byteArr = append(byteArr, bytes...) // storing bytes
	}
	return byteArr, nil
}

/**
this func convert get an array of byte and convert it to bit
*/
func byteToBit(bytes []byte) (string,error) {
	if len(bytes) <= 0{ //error handling
		err := fmt.Errorf("string is empty!")
		return "", err
	}

	bin := "" // empty string for storing bits
	for _, n := range bytes {
		bin = fmt.Sprintf("%s%.8b", bin, n) // storing bits
	}
	return bin, nil
}

/**
this func get an interface and reverse it
*/
func reverseStringArr(strArr []string) ([]string, error) {
	if len(strArr) <= 0 { //error handling
		err := fmt.Errorf("string is empty!")
		return nil, err
	}
	n := reflect.ValueOf(strArr).Len()
	swap := reflect.Swapper(strArr)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return strArr, nil
}

// ReverseString arg(Str) : is a string that you want to reverse it
// ReverseString return : is a reverse string that get form input of func
// ReverseString : that get a string and reverse all of that /**
func ReverseString(str string) (string, error) {
	if len(str) <= 0 { // error handling
		err := fmt.Errorf("string is empty!")
		return "", err
	}
	rns := []rune(str)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns), nil
}

// SeparatorString is a func that get a string and convert it to an array of 8 characters/**
func SeparatorString(str string) ([]string, error)  {
	if len(str) <= 0 { // error handling
		err := fmt.Errorf("string is empty!")
		return nil, err
	}
	revS , err := ReverseString(str)
	if err != nil { // error handling
		fmt.Errorf("you have a err: ", err)
	}

	rns := []rune(revS) // for convert string that reversed using "ReverseString" func to character
	var arr []string // for storing 8bits
	var zero = '0' // for add to bits if they are less than 8bits
	for len(rns)%8 != 0 {
		rns = append(rns, zero) // adding "zero" to bits if the length of last one is not 8
	}
	for i := 0; i < len(rns)/8 ; i++ {
		arr = append(arr, string(rns[8*i:8*i+8])) // storing 8bits
	}
	return arr, nil
}

// SeparatorStringArr is a func that get a string array and convert it to an array of 8 characters/**
func SeparatorStringArr(strArr []string) ([]string, error)  {
	if len(strArr) <= 0{ //error handling
		err := fmt.Errorf("string is empty!")
		return nil, err
	}
	revS, err := reverseStringArr(strArr)
	if err != nil {
		fmt.Errorf("you have err:", err)
	}
	zero := "0" // for add to bits if they are less than 8
	var arr []string // for storing 8bits
	for len(revS)%8 != 0  {
		revS = append(revS, zero) // adding "zero" to bits if the length of last one is not 8
	}
	for i := 0; i < len(revS)/8; i++ {
		arr = append(arr,revS[i*8] + revS[i*8+1] + revS[i*8+2]+ revS[i*8+3] + revS[i*8+4]+ revS[i*8+5] + revS[i*8+6]+ revS[i*8+7]) //storing 8bits
	}
	return arr, nil
}

// StringToBit is a func that get a string and convert it to a string of bits
func StringToBit(str string) (string, error) {
	if len(str) <= 0 { //error handling
		err := fmt.Errorf("sring is empty!")
		return "", err
	}

	chars, err1 := strToChar(str)
	byteArr, err2 := charToByte(chars)
	if err1 != nil { // error handling
		fmt.Errorf("you have an err: ", err1)
	}
	if err2 != nil { // error handling
		fmt.Errorf("you have an err: ", err2)
	}
	return byteToBit(byteArr)
}