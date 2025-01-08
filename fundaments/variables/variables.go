package main 

import "fmt"

var message = "This is a normal message"

func main(){
	var gretting = "Hello from the function main"
	fmt.Println(gretting)

	fmt.Println(message)

	message = "This another message"
	fmt.Println(message)

	fmt.Println("Esc for out")
}