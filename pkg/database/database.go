package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/christo-andrew/maybe-go/pkg/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func DefaultDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     os.Getenv("MYSQL_DATABASE_HOST"),
		Port:     os.Getenv("MYSQL_DATABASE_PORT"),
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		DBName:   os.Getenv("MYSQL_DATABASE_NAME"),
	}
}

func (config DatabaseConfig) ConnectionString() string {
	return config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DBName + "?parseTime=true"
}

func GetDB(databaseConfig DatabaseConfig) *gorm.DB {
	if databaseConfig == (DatabaseConfig{}) {
		databaseConfig = DefaultDatabaseConfig()
	}

	db, err := gorm.Open(mysql.Open(databaseConfig.ConnectionString()), &gorm.Config{
		Logger: logging.DatabaseQueryLogger(),
	})

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
