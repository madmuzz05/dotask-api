package initialize

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectToDB() *gorm.DB {
	LoadInitializeEnv()
	var err error
	dsn := os.Getenv("DB_URL")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "dotask.",
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}
