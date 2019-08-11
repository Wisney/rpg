package email

import (
	"fmt"
	"io/ioutil"
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
	gmail := gomail.NewDialer("smtp.gmail.com", 587, "thenoobsrpg@gmail.com", "porcos00")
	//gmail.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return gmail
}
