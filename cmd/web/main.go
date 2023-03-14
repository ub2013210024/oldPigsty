//

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

// Create a new type
type application struct {
	//question models.QuestionModel
}

func main() {
	// Create a flag for specifying the port number for when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("NAME_OF_DB"), "PostgreSQL dsn")
	flag.Parse()

	// Create an instance of the connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}

	// Create an instance of the application type
	app := &application{
		//question: models.QuestionModel{DB: db},
	}

	defer db.Close()
	log.Println("database connection pool established")

	// Create a customized server
	srv := &http.Server{
		Addr: *addr,
		//Handler: app.routes(),
	}
	log.Printf("starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
