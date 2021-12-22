/**
This package make a pattern of bits in 20*20 square with this pattern for each line-> array of 8 bits- reverse array of 8 bits - nand and array of 8 bits
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
	"errors"
	"fmt"
)

// PatternMaker param(str): is a string of our result encryption
// PatternMaker that get a string and make a pattern that is a 20*20 square of bits /**
func PatternMaker (str string) error {
	if len(str) <= 0 {
		return errors.New("string is empty!")
	}else {
		bits := bitSeparator.StringToBit(str)// we convert string to bits
		arr := bitSeparator.SeparatorString(bits) //we convert bits to 8 bits array
		var revArr []string //for holding reverse strings
		var nand string //for holding nand string
		var errBool bool //for error handling
		var nandArr []string // for holding nand strings from "bitwise package"

		//this loop reverse 8 bits array with help of "bitSeparator package"
		for i := 0; i < len(arr); i++ {
			revStr := bitSeparator.ReverseString(arr[i])
			revArr = append(revArr, revStr)
		}

		//this loop is for "nand and" with help of "bitWise package"
		for i := 0; i < len(arr); i++ {
			nand, errBool = bitWise.MakeNandAndCoupleBits(arr[i])
			if errBool {
				fmt.Println("you have a err during logic gate!")
			}
			nandArr = append(nandArr, nand)
		}

		//this loop is for printing our 20*20 square
		for i := 0; i <len(revArr) ; i++ {
			fmt.Println(arr[i] + revArr[i] + nandArr[i])
		}
	}
	return nil
}
