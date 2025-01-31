package activity
import "fmt"

func fibonacci(n int) int {
   if n <= 1 {
       return n
   }
   return fibonacci(n-1) + fibonacci(n-2)
}

func Fibo_sequence(n int) {
   fmt.Println("Fibonacci Numbers:")
   for i := 0; i < n; i++ {
       fmt.Printf("%v\t",fibonacci(i))
   }
}
