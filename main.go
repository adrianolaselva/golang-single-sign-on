package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/subchen/go-log"
	"github.com/subchen/go-log/formatters"
	_ "gopkg.in/go-oauth2/mysql.v3"
	"oauth2/src"
)

func main() {
	log.Default.Formatter = new(formatters.TextFormatter)
	app := src.Bootstrap{}
	app.Run()
}
