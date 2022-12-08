package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/kelompok4-loyaltypointagent/backend/config"
	"github.com/mailjet/mailjet-apiv3-go"
)

func ParseTemplate(filename string, data any) (string, error) {
	template, err := template.ParseFiles(filename)
	if err != nil {
		return "", err
	}

	buff := new(bytes.Buffer)
	if err = template.Execute(buff, data); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func Mailjet(to, subject, body string) error {
	mailjetConfig := config.LoadMailjetConfig()
	mailjetClient := mailjet.NewMailjetClient(mailjetConfig.APIKey, mailjetConfig.SecretKey)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: mailjetConfig.SenderEmail,
				Name:  mailjetConfig.SenderName,
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
				},
			},
			Subject:  subject,
			HTMLPart: body,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		return err
	}
	fmt.Printf("Email sent: %+v\n", res)

	return nil
}

func SendOTP(to string, data OTPEmailData) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	body, err := ParseTemplate(path.Join(wd, "html", "email", "otp_verification.html"), data)
	if err != nil {
		return err
	}

	if err := Mailjet(to, "OTP Verification", body); err != nil {
		return err
	}

	return nil
}

func SendAccessKey(to string, data AccessKeyEmailData) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	body, err := ParseTemplate(path.Join(wd, "html", "email", "forgot_password.html"), data)
	if err != nil {
		return err
	}

	if err := Mailjet(to, "Change your password!", body); err != nil {
		return err
	}

	return nil
}
