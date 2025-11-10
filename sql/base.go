package sql

import (
	"booksmanage/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// "root:QWEzxc123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local""
	sqlcfg := config.GlobeCfg.Mysql
	//dsn := sqlcfg.User + sqlcfg.Pwd + "@tcp(" + sqlcfg.IP + ":" + sqlcfg.Port + "/" + sqlcfg.DB + "?charset=utf8mb4&parseTime=True&loc=Local\""
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sqlcfg.User, sqlcfg.Pwd, sqlcfg.IP, sqlcfg.Port, sqlcfg.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("can't open sql, met err : %v", err))
	}
	DB = db
}
