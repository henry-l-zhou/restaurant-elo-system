package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// DBConn is a database connection instance
var DB *sql.DB

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
	if DB != nil {
		return DB
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
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Connected to Database")
	return DB
}

// WriteJSONData writes JSON data to a database table
func WriteJSONData(data []interface{}, table string, columns string) error {
	// Construct the query string
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (", table, columns)
	placeholders := make([]string, len(data))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	query += strings.Join(placeholders, ",") + ")"

	// Print the query and data for debugging purposes
	fmt.Println("Query: ", query)
	fmt.Println("Data: ", data)

	// Execute the query
	_, err := DB.Exec(query, data...)
	if err != nil {
		return err
	}

	return nil
}
