package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redmejia/api/handlers"
	"github.com/redmejia/api/router"
	"github.com/redmejia/internal/database"
)

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime)

	infoDbLog := log.New(os.Stdout, "INFO-DB\t", log.Ldate|log.Ltime)
	errorDbLog := log.New(os.Stdout, "ERROR-DB\t", log.Ldate|log.Ltime)

	greeting, db, err := database.DSNConnection(os.Getenv("DSN"))
	if err != nil {
		log.Println("ERROR CONN ", err)
	}

	infoDbLog.Println("Database created : ", greeting)

	defer db.Close()

	app := &handlers.App{
		Port:         ":80", // 80 on container
		ReadTimeOut:  5 * time.Second,
		WriteTimeOut: 10 * time.Second,
		ErrorLog:     errorLog,
		InfoLog:      infoLog,
		Db: &database.DbModel{
			DB:       db,
			InfoLog:  infoDbLog,
			ErrorLog: errorDbLog,
		},
	}

	app.Db.CreatingTestData()

	srv := &http.Server{
		Addr:         app.Port,
		Handler:      router.Router(app),
		ReadTimeout:  app.ReadTimeOut,
		WriteTimeout: app.WriteTimeOut,
	}
	infoLog.Println("Server is running at http://localhost:8080/v1")
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
