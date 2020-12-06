package libs

import (
	"log"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	config "github.com/spf13/viper"
)

var DB *gorm.DB

type DbConfig struct {
	Host         string
	Port         string
	Database     string
	User         string
	Password     string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
	TimeZone     string
	print_log    bool
}

func Configure(ConfigPath string, ConfigName string) DbConfig {
	config.AddConfigPath(ConfigPath)
	config.SetConfigName(ConfigName)
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if config.GetString("default.host") == "" {
		fmt.Println("falta host en el archivo " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.database") == "" {
		fmt.Println("falta database en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.user") == "" {
		fmt.Println("falta user en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	if config.GetString("default.password") == "" {
		fmt.Println("falta password en el archivo de config " + ConfigName)
		os.Exit(1)
	}
	
	response := DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
		"",
		config.GetBool("default.sql_log"),
	}


	return response
}
