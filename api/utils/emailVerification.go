package utils

import (
	"github.com/koybigino/food-app/api/email"
)

func SendEmail(token, mail, username string) {
	email.SendEmail("./public/src/emailTemplate.html", token, username, mail)
}
