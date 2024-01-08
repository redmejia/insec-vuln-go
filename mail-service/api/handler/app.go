package handler

import (
	"log"
	"time"

	"gopkg.in/gomail.v2"
)

type App struct {
	Port         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	ErrorLog     *log.Logger
	InfoLog      *log.Logger
	Mailer       *gomail.Message
	MailUser     string // stmp user uthenticated on server
	MailPassword string // stmp password uthenticad on server
	SmtpHost     string
	SmtpPort     int
}
