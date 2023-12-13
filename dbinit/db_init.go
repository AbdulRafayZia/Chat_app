package dbinit

import (
	"database/sql"
	"fmt"
	"log"
)

func Init() {
	// Connect to PostgreSQL server to run queries before database creation
	connInif := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)

	tempDB, err := sql.Open("postgres", connInif)
	if err != nil {
		log.Fatal(err)
	}

	defer tempDB.Close()

	// Run your queries here
	_, err = tempDB.Exec("CREATE EXTENSION IF NOT EXISTS " + dbname + ";")
	if err != nil {
		log.Fatal(err)
	}

	// Other queries as needed

	// Assign the temporary DB to the global variable for later use

}
