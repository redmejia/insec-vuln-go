package handler

import (
	"crypto/tls"
	"encoding/json"
	"net/http"

	"github.com/redmejia/mail/models"
	"gopkg.in/gomail.v2"
)

func (a *App) MailHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		var mail models.Mail
		err := json.NewDecoder(r.Body).Decode(&mail)
		if err != nil {
			a.ErrorLog.Fatal(err)
			return
		}

		a.sendMail(&mail)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(&mail)
		if err != nil {
			a.ErrorLog.Fatal(err)
		}
	}
}

func (a *App) sendMail(message *models.Mail) {

	a.Mailer.SetHeader("From", message.From)
	a.Mailer.SetHeader("To", message.To)
	a.Mailer.SetHeader("Subject", message.Subject)
	a.Mailer.SetBody("text/html", message.Body)

	dialer := gomail.NewDialer(a.SmtpHost, a.SmtpPort, a.MailUser, a.MailPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // skip verification

	if err := dialer.DialAndSend(a.Mailer); err != nil {
		a.ErrorLog.Fatal("unable to send email : ", err)
	}

	a.InfoLog.Println("message was sent")

}
