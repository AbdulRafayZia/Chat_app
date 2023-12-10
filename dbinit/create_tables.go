package dbinit

func CreateTables() error {

	db := OpenDB()

	defer db.Close()

	// Create tables if they don't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
				id SERIAL PRIMARY KEY,
				username VARCHAR(50) UNIQUE NOT NULL,
				password VARCHAR(100) NOT NULL,
				role VARCHAR(20) NOT NULL DEFAULT 'user'
		)
	`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS file_processes (
			id SERIAL PRIMARY KEY,
			filename VARCHAR(255) NOT NULL,
			words INTEGER,
			lines INTEGER,
			punctuations INTEGER,
			vowels INTEGER,
			execution_time FLOAT,
			routines INTEGER, 
			username VARCHAR(50)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
