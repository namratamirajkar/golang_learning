package input 
import (
	"fmt"
)

func GetInput() (value string){
	fmt.Print("Enter the input: ")
	fmt.Scan(&value)
	return value
}
