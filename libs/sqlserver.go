package libs

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func (c *DbConfig) InitSqlServerDB() *gorm.DB {
	// github.com/denisenkom/go-mssqldb
	//connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", c.User, c.Password, c.Host, c.Port, c.Database)
	//db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", c.Host, c.User, c.Password, c.Port, c.Database)
	db, err := gorm.Open("mssql", connectionString)

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
