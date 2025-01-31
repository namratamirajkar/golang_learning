package main
import (
	"fmt"
	"strconv"
	"example.com/go-activity-01/activity"
	"example.com/go-activity-01/input"
)

func main(){
	for{
		fmt.Println("\nChoose the Program to run \n 1. Reverse String \n 2. Fibonacci \n 3. Exit\n")
		fmt.Print("Enter your choice: ")
		var program string
		fmt.Scan(&program)
		input := input.GetInput()
		if program == "1"{
			activity.ReverseString(input)
			continue
		}else if program == "2"{
			n, err := strconv.Atoi(input)
			if err != nil{
				fmt.Println(err)
			}
			activity.Fibo_sequence(n)
			continue
		}else{
			return
		}
	}
}