package db

import (
	"net/url"
	"os"
	"crypto/tls"
)

import pg "github.com/go-pg/pg"

//GetConnect returns the connection of db
func GetConnect() *pg.DB {
	parsedURL, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	//parsedURL += "?sslmode=prefer"
	
	tlsConfig := &tls.Config{
        InsecureSkipVerify: true,
        // ServerName:         "localhost",
    }

	pgOptions := &pg.Options{
		User:     parsedURL.User.Username(),
		Database: parsedURL.Path[1:],
		Addr:     parsedURL.Host,
		TLSConfig: tlsConfig,
	}

	if password, ok := parsedURL.User.Password(); ok {
		pgOptions.Password = password
	}

	return pg.Connect(pgOptions)
}
