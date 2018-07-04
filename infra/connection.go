package infra

import (
	"net/url"
	"os"
)

import pg "github.com/go-pg/pg"

//GetConnect returns the connection of db
func GetConnect() *pg.DB {
	parsedURL, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	pgOptions := &pg.Options{
		User:     parsedURL.User.Username(),
		Database: parsedURL.Path[1:],
		Addr:     parsedURL.Host,
	}

	if password, ok := parsedURL.User.Password(); ok {
		pgOptions.Password = password
	}

	return pg.Connect(pgOptions)
}
