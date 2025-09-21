package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

func Migrations() {
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS migrations (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            filename TEXT UNIQUE NOT NULL,
            executed_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );
    `)
	if err != nil {
		log.Fatalf("Failed to create migrations table: %v", err)
	}

	files, err := filepath.Glob("internal/migrations/*.sql")
	if err != nil {
		log.Fatalf("Error reading migrations directory: %v", err)
	}
	sort.Strings(files)

	for _, file := range files {
		name := filepath.Base(file)
		var count int
		err := DB.QueryRow(
			"SELECT COUNT(*) FROM migrations WHERE filename = ?",
			name,
		).Scan(&count)
		if err != nil {
			log.Fatalf("Error checking migration %s: %v", name, err)
		}
		if count > 0 {
			continue
		}
		sqlBytes, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("Could not read migration file %s: %v", name, err)
		}
		_, err = DB.Exec(string(sqlBytes))
		if err != nil {
			log.Fatalf("Error executing migration %s: %v", name, err)
		}

		_, err = DB.Exec(
			"INSERT INTO migrations (filename) VALUES (?)",
			name,
		)
		if err != nil {
			log.Fatalf("Failed to record migration %s: %v", name, err)
		}

		fmt.Printf("Migration %s applied successfully\n", name)
	}
}
