package email

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

func SendEmail(templateEmailPath, token, username, toEmail string) {

	var body bytes.Buffer

	if t, err := template.ParseFiles(templateEmailPath); err != nil {
		fmt.Println(err)
		return
	} else {
		t.Execute(&body, struct {
			Token    string
			Username string
			AppName  string
		}{Token: token, Username: username, AppName: os.Getenv("APP_NAME")})
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("User_EMAIL"))
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Email Verification!")
	m.SetBody("text/html", body.String())
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("User_EMAIL"), os.Getenv("APP_KEY"))

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
