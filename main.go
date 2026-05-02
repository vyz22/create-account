package main

import (
	"fmt"
	"create-account/email"
)

func main() {
	mail := email.GenerateEmail()
	fmt.Println("Email:", mail)
}
