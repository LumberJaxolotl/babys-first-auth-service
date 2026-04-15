package dbhelpers

import (
	"fmt"
	"log"
	"os"
	"time"

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

func StoreRefreshToken(userId string, tokenStr string) {
	
	// calculating expires_at value
	const tokenExpiresIn = time.Hour * 24 * 60 
	expiresAt := time.Now().Add(tokenExpiresIn).Format(time.RFC3339)  
    tx := db.MustBegin()
    tx.MustExec(
		"INSERT INTO auth.refresh_tokens (user_id, token_hash, expires_at) VALUES ($1, $2, $3)", 
		userId, tokenStr, expiresAt)
    tx.Commit()
}

type RefreshToken struct {
	ID        string `db:"id"`
	user_id   string
	TokenHash string `db:"token_hash"`
}

func GetRefreshToken(tokenStr string) (RefreshToken, error) {
	// TODO validate this is the right method + sql code
	

	token := RefreshToken{}
	err := db.Get(&token, `
		SELECT id, user_id, token_hash
		FROM auth.refresh_tokens
		WHERE
			user_id = $1
			AND
			token_hash = $2
	`, tokenStr)
	if err != nil {
		log.Println("Token validation failed:", err)
		return RefreshToken{}, err
	}

	return token, nil
}



