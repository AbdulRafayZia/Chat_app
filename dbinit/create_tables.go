package dbinit

func CreateTables() error {

	db := OpenDB()

	defer db.Close()

	// Create tables if they don't exist
	_, err := db.Exec(`
		CREATE TABLE  users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(50) UNIQUE NOT NULL,
				password VARCHAR(100) NOT NULL,
				role VARCHAR(20) NOT NULL DEFAULT 'user'
		)
	`)
	if err != nil {
		return err
	}
	
	return nil
}
