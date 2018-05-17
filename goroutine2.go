package main

import ("fmt";
		"time")

func myFunc() {
    for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 100)
    }
}

func main() {
    go myFunc() // myFunc will run concurrently to the main implicit goroutine
  
    // Cheat and wait for user input to wait for the completion of the concurrent goroutine
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Go routine example")
}