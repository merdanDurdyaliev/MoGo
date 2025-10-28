package service

import (
	"net/smtp"
)

func SendLogin(mail, password, sendTo, msgs string) error {
	from := mail
	pass := password
	to := sendTo
   
	msg := "From: " + from + "\n" +
	"To: " + to + "\n" +
	"Subject: Hello there! You logged in!\n\n" +
	msgs
   
	err := smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	from, []string{to}, []byte(msg))
   
	if err != nil {
	 return err
	}
	
	
	return nil
}

func SendSignUp(mail, password, sendTo, Username, msgs string) error {
	who := Username
	from := mail
	pass := password
	to := sendTo
   
	msg := "Who " + who + "From: " + from + "\n" +
	"To: " + to + "\n" +
	"Subject: Hello there! You logged in!\n\n" +
	msgs
   
	err := smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	from, []string{to}, []byte(msg))
   
	if err != nil {
	 return err
	}
	
	return nil
}


