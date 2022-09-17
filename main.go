package main

import (
	"bufio"
	"chat-client/messages"
	"chat-client/users"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	go monitorReceiving()
	monitorSending()
}

func monitorReceiving(){
	for{
		msgs := messages.GetUnreadMessages("a8a5ba84-a113-4c29-aec5-4f73cf7a0cb5")
		for _, v := range msgs{
			println(messages.GetMessagePaload(v.Payload))
		}
		time.Sleep(5000)
	}
}

func initChat(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	fmt.Print("-> ")
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "", -1)
	fmt.Println("What is your email?")
	fmt.Print("-> ")
	email, _ := reader.ReadString('\n')
	email = strings.Replace(email, "\n", "", -1)
	users.CreateUser(name, email)
	fmt.Println("User created!")
}

func monitorSending(){
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("What is the message?")
		fmt.Print("-> ")
		message, _ := reader.ReadString('\n')
		message = strings.Replace(message, "\n", "", -1)

		messages.SendMessage("a8a5ba84-a113-4c29-aec5-4f73cf7a0cb5", message)
	}
}
