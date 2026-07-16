package database

import (
	"database/sql"
    "os"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	var err error

    log.Println("1 - DATABASE_URL:", os.Getenv("DATABASE_URL"))

    DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        return err
    }

    log.Println("2 - sql.Open executado")
    log.Println("3 - DB:", DB)

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("✅ Banco conectado!")

	return createTables()
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS simulations (
		id SERIAL PRIMARY KEY,

		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

		current_assets DOUBLE PRECISION NOT NULL,
		monthly_contribution DOUBLE PRECISION NOT NULL,
		annual_percentage DOUBLE PRECISION NOT NULL,
		current_age INTEGER NOT NULL,
		retirement_age INTEGER NOT NULL,
		time_in_years INTEGER,
		inflation DOUBLE PRECISION NOT NULL,

		final_amount DOUBLE PRECISION NOT NULL,
		inflation_adjusted_amount DOUBLE PRECISION NOT NULL,
		years_to_retirement INTEGER NOT NULL,
		real_rate_year DOUBLE PRECISION NOT NULL,
		real_rate_month DOUBLE PRECISION NOT NULL,
		total_contributed DOUBLE PRECISION NOT NULL,
		total_interest_earned DOUBLE PRECISION NOT NULL
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
