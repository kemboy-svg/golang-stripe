package Config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"strings"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	// Split the EliasLocal string to extract credentials
	credentials := strings.SplitN(EliasLocal, "@", 2)
	userPass := strings.SplitN(credentials[0], ":", 2)
	
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     userPass[0],
		Password: userPass[1],
		DBName:   "products",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

const EliasLocal string = "elias:elias123@tcp(127.0.0.1:3306)/products?parseTime=true"
