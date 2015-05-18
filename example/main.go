package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/thoas/djangogo/example/application"
	"os"
)

func main() {
	app, err := application.New(os.Getenv("COOKIE_NAME"), &application.Option{
		Database: map[string]string{
			"USER":     os.Getenv("DATABASE_USER"),
			"NAME":     os.Getenv("DATABASE_NAME"),
			"PASSWORD": os.Getenv("DATABASE_PASSWORD"),
		},
		Session: map[string]string{
			"PREFIX":   os.Getenv("SESSION_PREFIX"),
			"PORT":     os.Getenv("REDIS_PORT"),
			"DATABASE": os.Getenv("REDIS_DATABASE"),
		},
	})

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	app.Run(3001)
}
