package connection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"joke/conf"
)

var (
	DB *gorm.DB
)

func Connect(config *conf.MysqlConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("db error: %v", err)
	}
}
