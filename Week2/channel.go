package main

import (
	"fmt"
	"time"
)

func main1() {
	channel := make(chan string, 1)
	go func(ch chan<- string) {
		ch <- "2"
		fmt.Println(1)
	}(channel)
	message := <-channel
	time.Sleep(time.Second * 5)
	fmt.Println(message)
}

// func main() {
// 	channel := make(chan string, 1)
// 	go func(ch <-chan string) {
// 		// ch <- "2"
// 		mess := ch
// 		fmt.Println(mess)
// 		fmt.Println(1)
// 	}(channel)
// 	message := "Hello frm Main Function"
// 	channel <- message
// 	fmt.Println(message)
// }

func main2() {
	channel := make(chan string, 1)
	go func(ch <-chan string) {

		mess := <-ch
		fmt.Println(mess)
		fmt.Println(1)
	}(channel)
	message := "Hello frm Main Function"
	channel <- message
	time.Sleep(time.Second * 5)
	fmt.Println("message")
}
