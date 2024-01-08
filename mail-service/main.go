package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/redmejia/mail/api/handler"
	"github.com/redmejia/mail/api/router"
	"gopkg.in/gomail.v2"
)

func main() {

	message := gomail.NewMessage()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	// SmtpHost:     "smtp.katlinaxample.com", // use mailhog host
	// SmtpPort:     587,                      // 1025 port mailhog
	stmpPort, _ := strconv.Atoi(os.Getenv("port"))
	app := &handler.App{
		Port:         ":587", // container smtp server
		ReadTimeOut:  5 * time.Second,
		WriteTimeOut: 10 * time.Second,
		ErrorLog:     errorLog,
		InfoLog:      infoLog,
		Mailer:       message,
		MailUser:     "kat@test.me",
		MailPassword: "TheFlowerAreAzulComoElsky",
		SmtpHost:     os.Getenv("host"), // use mailhog host
		SmtpPort:     stmpPort,          // 1025 port mailhog
	}

	srv := &http.Server{
		Addr:         app.Port,
		Handler:      router.Router(app),
		ReadTimeout:  app.ReadTimeOut,
		WriteTimeout: app.WriteTimeOut,
	}

	infoLog.Println("Server is running at http://localhost:8081/v1")
	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}

}
