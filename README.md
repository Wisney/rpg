Site:
```
    https://rpgzinho.herokuapp.com/
```

To install and run:

In Postgres:
```
    CREATE EXTENSION hstore;
```

In Project:
```
go get github.com/go-pg/pg

go get gopkg.in/gomail.v2

go get github.com/dgrijalva/jwt-go

go get github.com/gorilla/mux

```

To init:
```
go install ./...

heroku local
```
