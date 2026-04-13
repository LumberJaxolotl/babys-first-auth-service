package dbhelpers

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	// Format: postgres://username:password@localhost:5432/database_name
	dsn := "postgres://app_user:password@localhost:5432/app_db"

	// 1. Create a connection
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db = conn

	// 2. Test the connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("Could not ping DB:", err)
	}

	// db connection never gets closed ever in this example, 
	// but thats ok for an example:)
	fmt.Println("Successfully connected to Postgres!")
}

// TODO build this out 
func StoreRefreshToken() {
	db.MustExec(`INSERT VALUES FROM User`)
}

func ValidateRefreshToken(tokenStr string) {
	// TODO validate this is the right method + sql code
	db.Select(`
		SELECT user_id, token_hash 
		FROM User
		WHERE 
			user_id = 
			AND
			token_hash = 
	`)
}



