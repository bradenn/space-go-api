package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sql"
	"net/url"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	env := GetConfig()
	dbConfig := DBConfig{
		Host:     env.DBHost,
		Port:     env.DBPort,
		User:     env.DBUser,
		DBName:   env.DBName,
		Password: env.DBPass,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	u := url.URL{
		User:     url.UserPassword(dbConfig.User, dbConfig.Password),
		Scheme:   "mysql",
		Host:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
		Path:     dbConfig.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"enabled"}}).Encode(),
	}
	return  u.String()
}
