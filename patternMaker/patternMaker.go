/**
This package make a pattern of bits in 20*20 square
This package use "bitSeparator package" for convert string to char and then convert it to byte and then convert it to bits
This package use "bitWise package" for nand bits
*/

package patternMaker

/**
@Auth: Mohammad Rezaie
@Date: 19 Dec 2021
*/


/**
******NOTE: before using this package make sure that you import "bitSeparator" & "bitWise" packages
*/

import (
	//"Your module/address/Your bitSeparator package
	//"Your module/address/Your bitWise package
	"fmt"
)

// PatternMaker that get a string and make a pattern that is a 20*20 square of bits /**
func PatternMaker (s string) {
	bits := bitSeparator.StringToBit(s)
	arr := bitSeparator.SeparatorString(bits)
	var revArr []string
	var nand string
	var errBool bool
	var nandArr []string
	for i := 0; i < len(arr); i++ {
		v := bitSeparator.ReverseString(arr[i])
		revArr = append(revArr, v)

	}

	for i := 0; i < len(arr); i++ {
		nand, errBool = bitWise.MakeNandAndCoupleBits(arr[i])
		if errBool {
			fmt.Println("you have a err during logic gate!")
		}
		nandArr = append(nandArr, nand)

	}

	for i := 0; i <len(revArr) ; i++ {
		fmt.Println(arr[i] + revArr[i] + nandArr[i])
	}
}

