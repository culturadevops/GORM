package libs

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (c *DbConfig) InitMysqlDB() *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	db.SingularTable(true)                  //全局设置表名不可以为复数形式。
	db.DB().SetMaxIdleConns(c.MaxIdleConns) //空闲时最大的连接数
	db.DB().SetMaxOpenConns(c.MaxOpenConns) //最大的连接数
	return db
}
