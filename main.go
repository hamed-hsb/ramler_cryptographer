package main 



import(
"encrypter.com/ramler/sha256hash"
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
		
		onSelector()
		
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


	func onSelector(){

		selector := 0
		data := ""
		
		Println("Enter number item of hash algorithm(as 1 to 6) :")
		Scanf("%d", &selector)
		
		switch selector {
			case 1:
				Println("MD2 algorithm selected .")
			case 2:
				Println("MD4 algorithm selected .")
			case 3:
				Println("MD5 algorithm selected .")
			case 4:
				Println("MD6 algorithm selected .")
			case 5:
				Println("SHA256 algorithm selected .")
			case 6:
				Println("BSOGL algorithm selected .")
			
		}
		
		
		Println("Write anything :")
		Scanf("%s",&data)
		Println("in --> ",data)
		switch selector {
			case 1:
				
			case 2:
				
			case 3:
				
			case 4:
				
			case 5:
				shaHashCode(data)
				Println("\n")
				Println("\n")
			case 6:
				
			
		}
		
		
		
			
	}

	func onHelp(){
		Println("\n")
		Println("You can using of these hash algorithms: ")
		Println("1- MD2 hash algorithm and return 128 bits hash.")
		Println("2- MD4 hash algorithm and return 128 bits hash.")
		Println("3- MD5 hash algorithm and return 128 bits hash.")
		Println("4- MD6 hash algorithm and return up to 512 bits hash.")
		Println("5- SHA256 hash algorithm and return up to 256 Byte hash.")
		Println("6- BSOGL (bit security of gate login) this algorithm return a hash string.")
		Println("\n")
	}
	
	func shaHashCode(data string){
	
		hash := sha256.Hash(data)
		Println("out --> ",hash)
	}















