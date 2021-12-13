package main

import (
	"fmt"
	"ramler-golang/bitwise/package"
)

func main() {
	//bitwise.And(uint8(strconv.ParseInt("00000001", 2, 64)),uint8(strconv.ParseInt("00000001", 2, 8)))
	var a byte = 5
	var b byte = 3
	fmt.Println(bitwise.Or(a, b))
	fmt.Println(bitwise.Nor(a, b))
}
