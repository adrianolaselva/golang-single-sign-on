package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "gopkg.in/go-oauth2/mysql.v3"
	"oauth2/src"
)

func main() {
	app := src.Bootstrap{}
	app.Run()
}
