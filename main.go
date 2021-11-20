package main

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	fmt.Println("asd")

	mail := gomail.NewMessage()
	mail.SetHeader("From", "*******@gmail.com")
	mail.SetHeader("To", "*****@icloud.com")
	mail.SetHeader("Subject", "Respuesta de solicitud")
	mail.SetBody("text/plain", "Esto es una prueba para mandar emails con Golang.")

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "*********¨*@gmail.com", "HereGoesPassowrd")
	if err := dialer.DialAndSend(mail); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Email Enviado! ")

}
