package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

type Database struct {

}

func (c *Database) Connect() *gorm.DB {

	db, err := gorm.Open("mysql", os.Getenv("SSO_DB_CONNECTION_STRING"))
	if err != nil {
		log.Printf("failed to connect database: %s", err)
	}

	return db
}
