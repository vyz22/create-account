package main

import (
	"fmt"
	"strings"
	"time"

	"create-account/email"
	"create-account/otp"
)

func main() {
	mail := email.GenerateEmail()
	fmt.Println("Email:", mail)

	// pisahin login & domain
	parts := strings.Split(mail, "@")
	login := parts[0]
	domain := parts[1]

	fmt.Println("Menunggu OTP...")

	for i := 0; i < 12; i++ {
		msgs, _ := otp.GetMessages(login, domain)

		if len(msgs) > 0 {
			body, _ := otp.ReadMessage(login, domain, msgs[0].ID)
			code := otp.ExtractOTP(body)

			if code != "" {
				fmt.Println("OTP ditemukan:", code)
				return
			}
		}

		fmt.Println("Belum ada OTP...")
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Gagal ambil OTP")
}
