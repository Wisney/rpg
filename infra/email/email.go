package email

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

//SendEmail send a email .-.
func SendEmail(emailTo string, title string, text string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "thenoobsrpg@gmail.com")
	mail.SetHeader("To", emailTo)
	mail.SetHeader("Subject", title)
	mail.SetBody("text/html", text)

	d := getEmailManager()

	if err := d.DialAndSend(mail); err != nil {
		panic(err)
	}
}

//SendForgotPasswordEmail send a email .-.
func SendForgotPasswordEmail(emailTo string, token string) {
	forgotpassword, err := ioutil.ReadFile("./infra/email/email_forgot_password.html")
	if err != nil {
		fmt.Print(err)
	}

	title := "[The Noobs] Redefinir Senha"
	text := string(forgotpassword)
	text = strings.Replace(text, "{{.LinkWithToken}}", os.Getenv("SITE")+"/resetpassword/"+token, 1)

	SendEmail(emailTo, title, text)
}

func getEmailManager() *gomail.Dialer {
	mailman, _ := url.Parse(os.Getenv("MAILMAN"))
	password, _ := mailman.User.Password()
	mailDialer := gomail.NewDialer(mailman.Host, 587, mailman.User.Username(), password)
	//mailDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return mailDialer
}
