package infra

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetDB(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	categoryTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255)
	);`

	courseTable := `
	CREATE TABLE IF NOT EXISTS courses (
		id VARCHAR(255) PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255),
		category_id VARCHAR(255),
		FOREIGN KEY (category_id) REFERENCES categories(id)
	);`

	_, err := db.Exec(categoryTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(courseTable)
	if err != nil {
		return err
	}

	return nil
}
