package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ") 
	var input uint
	fmt.Scanf("%d", &input)
	output := fibonacci(input)
	fmt.Println("Result: ",output)
}

func fibonacci(x uint) uint{
	if x==0 {
		return 0
	} else if x==1 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}