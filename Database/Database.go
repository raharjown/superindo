package Database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	host     string
	port     string
	user     string
	dbname   string
	password string
	sslmode  string
}

func BuildConfig() *DBConfig {
	dbConfig := DBConfig{
		host:     os.Getenv("DB_host"),
		port:     os.Getenv("DB_port"),
		user:     os.Getenv("DB_user"),
		password: os.Getenv("DB_password"),
		dbname:   os.Getenv("DB_name"),
		sslmode:  os.Getenv("DB_sslmode"),
	}
	return &dbConfig
}

func DBURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.host,
		dbConfig.port,
		dbConfig.user,
		dbConfig.password,
		dbConfig.dbname,
		dbConfig.sslmode,
	)
}
