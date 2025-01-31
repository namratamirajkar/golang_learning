package activity

import(
	"fmt"
)

func ReverseString(str string){

	// coverting the string into character slice
		strbytes := []rune(str) // stores ASCII value of each char in string as element of array strbytes
	
		// i and j are short declaration
		//used multiple declaration and assignment of value to variable
		
		for i, j := 0, len(strbytes)-1; i < j ; i, j = i+1, j-1{
			 strbytes[i], strbytes[j] = strbytes[j], strbytes[i]
		}
		
		// string("string") coverts array of char to "string"
		
		fmt.Println(" The reverse string of given string is :", string(strbytes))
	
}




