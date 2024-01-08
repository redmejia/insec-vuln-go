package handlers

import (
	"log"
	"time"

	"github.com/redmejia/internal/database"
)

type App struct {
	Port         string
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration
	ErrorLog     *log.Logger
	InfoLog      *log.Logger
	Db           *database.DbModel
}
