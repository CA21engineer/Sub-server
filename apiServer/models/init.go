package models

import (
	"fmt"

	"github.com/BambooTuna/go-server-lib/config"

	"github.com/jinzhu/gorm"

	// mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB のグローバル変数
var DB *gorm.DB

// ConnectDB connect db
func ConnectDB() {
	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.GetEnvString("MYSQL_USER", "ca21engineer"),
		config.GetEnvString("MYSQL_PASS", "pass"),
		config.GetEnvString("MYSQL_HOST", "127.0.0.1"),
		config.GetEnvString("MYSQL_PORT", "3306"),
		config.GetEnvString("MYSQL_DATABASE", "subs"),
	)
	var err error
	DB, err = gorm.Open("mysql", connect)

	if err != nil {
		panic(err.Error())
	}
}
