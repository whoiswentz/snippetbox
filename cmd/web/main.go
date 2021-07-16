package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/whoiswentz/snippetbox/pkg/mysql"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	session *sessions.Session
	snippets *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dns", "root:root@/snippetbox?parseTime=true", "MySQL data source name")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret Key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "ERROR	\t", log.Ldate|log.Ltime)

	db, err := openDb(*dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := NewTemplateCache("./ui/html/")
	if err != nil {
		errLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := application{
		infoLog:  infoLog,
		errLog:   errLog,
		session: session,
		snippets: mysql.NewSnippetModel(db),
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s\n", *addr)
	if err := srv.ListenAndServe(); err != nil {
		errLog.Fatal(err)
	}
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
