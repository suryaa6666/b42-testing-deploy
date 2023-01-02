package mysql

import (
	"fmt"

	// Import postgres package here ...
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	// Get `host`, `user`, `password`, `database name`, and `port` from env here ...

	// ===== If using MySql =====
	// dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// ===== If using Postgres =====
	// Setup database connection here ...

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database")
}
