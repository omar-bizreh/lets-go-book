package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Web application
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	// Define a new command line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Create new logger
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// HTTP server
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on port: %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
