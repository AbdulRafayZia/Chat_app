package dbinit

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	host     = getEnv("DB_HOST", "localhost")
	port     = getEnvInt("DB_PORT", 5432)
	user     = getEnv("DB_USER", "postgres")
	password = getEnv("DB_PASSWORD", "1234567890")
	dbname   = getEnv("DB_NAME", "chat_app")
	connStr  = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
)

func OpenDB() *sql.DB {

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	// Create the database if it does not exist

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// var db *sql.DB

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if strValue == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		log.Printf("Error converting %s to integer: %v", key, err)
		return defaultValue
	}
	return value
}
