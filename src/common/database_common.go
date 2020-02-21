package common

import (
	"database/sql"
	"fmt"
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

	db.LogMode(true)

	return db
}

func (c *Database) InitializePaginate(conn *gorm.DB, out interface{}, filters *map[string]interface{}, orderBy *string, orderDir *string, limit int, page int, orderByDefault string, orderDirDefault string) (*sql.Rows, int, int, error) {
	var total int

	if page == 0 {
		page = 1
	}

	if len(*orderBy) == 0 {
		orderBy = &orderByDefault
	}

	if len(*orderDir) == 0 {
		orderDir = &orderDirDefault
	}

	filterDefault := make(map[string]interface{})
	if filters == nil {
		filters = &filterDefault
	}

	offset := (page-1)*limit
	order := fmt.Sprintf("%s %s", *orderBy, *orderDir)

	result := conn.
		Order(order).
		Where(*filters).
		Find(out)

	if result.Error != nil {
		return nil, 0, 0, result.Error
	}

	result.Count(&total)

	rows, err := result.
		Limit(limit).
		Offset(offset).
		Rows()

	if err != nil {
		return nil, 0, 0, err
	}

	pages := (total+limit)/limit

	return rows, total, pages, nil
}
