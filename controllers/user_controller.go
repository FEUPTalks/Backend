package controllers

import (
	"net/http"
	"net/smtp"

	"fmt"

	"strings"

	"github.com/FEUPTalks/Backend/model"
)

//UserController struct
type UserController struct {
}

// SendEmailValidation sent an email to user in order to get the validation link
func SendEmailValidation(user *model.User, request *http.Request) error {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"feuptalks@gmail.com",
		"Talks123",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	subject := "FEUPTalks Validation"
	from := "feuptalks@gmail.com"
	//body := ParseTemplateHTML(user.Name, fmt.Sprintf("http://%s%s/%s", request.Host, request.URL.Path, user.HashCode))
	body := ParseTemplate(user.Name, fmt.Sprintf("http://%s%s/%s", request.Host, request.URL.Path, user.HashCode))
	msg := "Subject: " + subject + "\n" +
		"From: " + from + "\n" +
		"To: " + user.Email + "\n" +
		body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"feuptalks@gmail.com",
		[]string{user.Email},
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
		
To access FEUPtalks you don't need any passoword. You just need to follow this link:
"STRURL"

FEUPTalks`)
	template = strings.Replace(template, "STRNAME", name, -1)
	template = strings.Replace(template, "STRURL", urlValid, -1)
	return template
}
