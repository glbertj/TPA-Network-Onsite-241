package utils

import (
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASS"), os.Getenv("MAIL_HOST"))
	msg := []byte(
		"From: stv <" + os.Getenv("MAIL_USER") + ">\n" +
			"To: " + to + "\n" +
			"Subject: " + subject + "\n" +
			"MIME-Version: 1.0\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\n\n" +
			body)

	return smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"), auth, os.Getenv("MAIL_USER"), []string{to}, msg)
}
