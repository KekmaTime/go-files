package main

import "fmt"

type messagetoSend struct {
	phoneNumber int64
	message string
}

func test(m messagetoSend) {
	fmt.Printf("Sending message: '%s' to : %v\n", m.message, m.phoneNumber)
	fmt.Println("================================================")
}

func main() {
	test(messagetoSend{
		phoneNumber: 148255510981,
		message: "You should run!",
	})
	test(messagetoSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messagetoSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
}