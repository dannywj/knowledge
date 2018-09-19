package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	fmt.Println("========begin task=========")
	m := gomail.NewMessage()
	m.SetHeader("From", "dannywj@live.cn")
	m.SetHeader("To", "dannywj@live.cn")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Alarm-Planting Queue Full!")
	m.SetBody("text/html", "Planting Tree Queue item:10")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.office365.com", 587, "dannywj@live.cn", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("send success")
	fmt.Println("========end task=========")

}
