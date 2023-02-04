package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DBConn is a database connection instance
var DBConn *sql.DB

// Config is the global database configuration
var Config DBConfig

// DBConfig contains configuration for connecting to the database
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Connect opens a database connection
func Connect(cfg ...DBConfig) *sql.DB {
	if DBConn != nil {
		return DBConn
	}
	if len(cfg) > 0 {
		Config = cfg[0]
	} else {
		Config = DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_HOST"),
		}
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		Config.Host,
		Config.Port,
		Config.User,
		Config.Password,
	)
	var err error
	DBConn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Connected to Database")
	return DBConn
}

// WriteJSONData writes JSON data to a database table
func WriteJSONData(data interface{}, table string) error {
	DBConn := Connect()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = DBConn.Exec("INSERT INTO "+table+" (data) VALUES ($1)", jsonData)
	if err != nil {
		return err
	}

	return nil
}
