package main

import (
	."fmt"
	. "ramler-golang/bitwise"
)



func main(){
	var input string
	Scanf("%s",&input)
	hashed , error := MakeNandAndCoupleBits(input)
	if(!error){
		println(hashed)
	}else{
		println("there is error"+ hashed)
	}
}



