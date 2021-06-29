package datastore

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mashita1023/clean-architecture/conf"
)

func NewMySqlDB() *gorm.DB {

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		conf.C.Database.User,
		conf.C.Database.Password,
		conf.C.Database.Host,
		conf.C.Database.Port,
		conf.C.Database.Dbname,
	)

	fmt.Println(connectionString)

	dbconf, err := gorm.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	err = dbconf.DB().Ping()
	if err != nil {
		panic(err)
	}

	dbconf.LogMode(true)

	dbconf.Set("gorm:table_options", "ENGINE=InnoDB")

	return dbconf
}
