/**
this package convert a string to bits and convert bits to array of 8 bits
 */

package bitSeparator

/**
@au Mohammad Rezaie
@date 13 Dec 2021
*/

import (
	"fmt"
	"reflect"
)

/**
this func get a string and convert it to an array of characters
*/
func strToChar(str string) []string {
	var charArr []string
	chars := []rune(str)
	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		charArr = append(charArr, char)

	}
	return charArr
}

/**
this func convert characters to an array of byte
*/
func charToByte(chars []string) []byte {
	var byteArr []byte
	for i := 0; i < len(chars); i++ {
		bytes := []byte(chars[i])
		byteArr = append(byteArr, bytes...)
	}
	return byteArr
}

/**
this func convert get a array of byte and convert it to bit
*/
func byteToBit(bytes []byte) string{
	bs := []byte(bytes)
	bin := ""
	for _, n := range bs {
		bin = fmt.Sprintf("%s%.8b", bin, n)
	}
	return bin
}


/**
this func get an interface and reverse it
*/
func reverseStringArr(s []string) []string {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return s
}

/**
this func get a string and reverse it
*/
func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)

}

// SeparatorString is a func that get a string and convert it to an array of 8 characters/**
func SeparatorString(s string) []string  {
	revS := reverseString(s)
	rns := []rune(revS)
	var arr []string
	var zero = '0'
	for len(rns)%8 != 0 {
		rns = append(rns, zero)
	}
	for i := 0; i < len(rns)/8 ; i++ {
		arr = append(arr, string(rns[8*i:8*i+8]))
	}
	return arr
}

// SeparatorStringArr is a func that get a string array and convert it to an array of 8 characters/**
func SeparatorStringArr(s []string) []string  {
	revS := reverseStringArr(s)
	zero := "0"
	var arr []string
	for len(revS)%8 != 0  {
		revS = append(revS, zero)
	}
	for i := 0; i < len(revS)/8; i++ {
		arr = append(arr,revS[i*8] + revS[i*8+1] + revS[i*8+2]+ revS[i*8+3] + revS[i*8+4]+ revS[i*8+5] + revS[i*8+6]+ revS[i*8+7])
	}
	return arr
}

// StringToBit is a func that get a string and convert it to a string of bits
func StringToBit(s string) string {
	chars := strToChar(s)
	byteArr := charToByte(chars)
	return byteToBit(byteArr)
}
