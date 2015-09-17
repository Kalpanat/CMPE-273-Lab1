package main

import ( "fmt"; "time")

func sleep(c time.Duration) {
	select { 
		case <- time.After(c):
		fmt.Println("timeout", c) 
		}
	}

func main() {
	fmt.Println("Hi")
	sleep(time.Second*3)
	fmt.Println("After Sleep print")
}