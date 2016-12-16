package services

import (
	"net/smtp"
	"strings"

	"github.com/FEUPTalks/Backend/model"
)

const subject string = "FEUPTalks Validation"
const from string = "feuptalks@gmail.com"

// SendEmailConfirmation send an email to user in order to get the validation link
func SendEmailConfirmation(email *model.Email) error {

	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"feuptalks@gmail.com",
		"Talks123",
		"smtp.gmail.com",
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	body := ParseTemplate(email.NameTo, email.URL)
	msg := "Subject: " + subject + "\n" +
		"From: " + from + "\n" +
		"To: " + email.EmailTo + "\n" +
		body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"feuptalks@gmail.com",
		[]string{email.EmailTo},
		[]byte(msg),
	)
	return err
}

//ParseTemplateHTML fill template html
func ParseTemplateHTML(name string, urlValid string) string {
	template :=
		(`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN""http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<html></head><body><p>Hello STRNAME
	<a href="STRURL">Validation address</a>
	</p></body></html>`)
	template = strings.Replace(template, "STRNAME", name, -1)
	template = strings.Replace(template, "STRURL", urlValid, -1)
	return template
}

//ParseTemplate fill template in Text
func ParseTemplate(name string, urlValid string) string {
	template := (`
Hello STRNAME, 
		
To access FEUPtalks you don't need any password. You just need to follow this link:

"STRURL"

FEUPTalks`)

	template = strings.Replace(template, "STRNAME", name, -1)
	template = strings.Replace(template, "STRURL", urlValid, -1)
	return template
}
