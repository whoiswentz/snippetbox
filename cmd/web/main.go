package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR	\t", log.Ldate|log.Ltime)

	app := application{
		infoLog: infoLog,
		errLog:  errLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errLog.Fatal(err)
	}
}
