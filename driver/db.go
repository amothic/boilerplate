package driver

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenSQL() (*sql.DB, error) {
	return sql.Open("mysql", "root@tcp(db:3306)/boilerplate?charset=utf8mb4&parseTime=True&loc=Local")
}
