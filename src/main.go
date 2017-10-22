package main

import "fmt"

func main() {

	const p int = 2

	request := Request{p: p, hasRequest: false}

	fmt.Println(request.p)

	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}
