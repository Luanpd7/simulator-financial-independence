package database

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func Init(dbPath string) error {
	var err error
	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	if err = createTables(); err != nil {
		return err
	}


	return nil
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS simulations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,

		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		current_assets REAL NOT NULL,
		monthly_contribution REAL NOT NULL,
		annual_percentage REAL NOT NULL,
		current_age INTEGER NOT NULL,
		retirement_age INTEGER NOT NULL,
		time_in_years INTEGER,
		inflation REAL NOT NULL,

		final_amount REAL NOT NULL,
		inflation_adjusted_amount REAL NOT NULL,
		years_to_retirement INTEGER NOT NULL,
		real_rate_year REAL NOT NULL,
		real_rate_month REAL NOT NULL,
		total_contributed REAL NOT NULL,
		total_interest_earned REAL NOT NULL
	);
	`

	_, err := DB.Exec(query)
	return err
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
