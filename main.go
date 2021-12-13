package main 



import(

	."fmt"

	_"strconv"

)





func main(){

	a := "y"
	mySelect := 0
	myStr := ""
	help := ""
	for a == "y" {
		a = "n"
		myStr = ""
		
		onStart()
		Scanf("%s", &help)
		
		if(help == "yes" || help == "y"){
			onHelp()
		}
		
		
		Println("1: select file \n2: enter string \n---------------------------\n--->")
		Scanf("%d", &mySelect)
		switch mySelect {
		case 1:
			Println("enter relative path of your file")
		case 2:
			Println("Enter you string")
			Scanf("%s", &myStr)
			/*Scanf("%s", &myStr)
			result := ""
			for char := range myStr{
				result += explode(int64(myStr[char]))
			}
			add0(&result)
			println("main", result)

			var hashed []uint8;
			for i:=0; i<len(result); i+=8{
				println("slices",result[i:i+8])
				tes, err := strconv.ParseInt(result[i:i+8], 2, 64)
				if err == nil {
					println("tes 64", tes)
					println("tes unit8",uint8(tes))
					hashed = append(hashed, uint8(tes))
				}else {
					println("error",err)
				}
			}
			println("hased",hashed)
			println("final hashed",string(hashed))*/
		}
		Println("Do you want to continue?(y/n)")
		Scanf("%s", &a)
	}

}


func onStart(){

Println("This is a program to generate a hash string. ramler can use different hash algorithms.")
Println("If you will help enter Yes: ")
Println("")
}


func onHelp(){
Println("\n")
Println("You can using of these hash algorithms: ")
Println("1- MD2 hash algorithm and return 128 bits hash.")
Println("2- MD4 hash algorithm and return 128 bits hash.")
Println("3- MD5 hash algorithm and return 128 bits hash.")
Println("4- MD6 hash algorithm and return up to 512 bits hash.")
Println("5- BSOGL (bit security of gate login) this algorithm return a hash string.")
Println("\n")
}















