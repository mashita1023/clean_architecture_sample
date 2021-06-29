package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mashita1023/clean-architecture/conf"
	"github.com/mashita1023/clean-architecture/infrastructure/datastore"
)

func main() {

	conf.ReadConf()

	dbconf := datastore.NewMySqlDB()
	fmt.Println(dbconf)
	r := gin.Default()

	r.Run(":" + conf.C.Server.Address)

}
