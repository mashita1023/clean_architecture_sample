package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}

	Server struct {
		Address string
	}
}

var C config

func ReadConf() {
	Conf := &C

	viper.SetConfigName("config")

	viper.SetConfigType("yml")

	//	viper.AddConfigPath(filepath.Join("conf"))
	// ここのパスを書く
	viper.AddConfigPath(filepath.Join("$GOPATH", "src", "github.com", "mashita1023", "clean-architecture", "conf"))

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file read error")
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println("config file Unmarshal error")
		fmt.Println(err)
		os.Exit(1)
	}
}
