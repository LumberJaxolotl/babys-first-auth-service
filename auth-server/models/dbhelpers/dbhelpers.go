package dbhelpers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// ---------------------- Postgres Boilerplate -------------------
var db *sqlx.DB
func init() {
	// Format: postgres://username:password@localhost:5432/database_name
	dsn := "postgres://app_user:password@localhost:5432/app_db?sslmode=disable"

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
// gets a connection to the db for convenience in testing 
func getDB()*sqlx.DB{
	return db
}

// ---------------------- END Postgres Boilerplate -----------------

// --------------------- Refresh Token CRUD ------------------------ 

type RefreshToken struct {
	ID        string `db:"id"`
	user_id   string
	TokenHash string `db:"token_hash"`
}

func StoreRefreshToken(userId string, tokenStr string) {
	
	// calculating expires_at value
	const tokenExpiresIn = time.Hour * 24 * 60 
	expiresAt := time.Now().Add(tokenExpiresIn).Format(time.RFC3339)  
    tx := db.MustBegin()
	tx.MustExec(`
		INSERT INTO auth.refresh_tokens 
		(user_id, token_hash, expires_at) 
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id) DO UPDATE SET
			token_hash = EXCLUDED.token_hash,
			expires_at = EXCLUDED.expires_at
	`, userId, tokenStr, expiresAt)
    tx.Commit()
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

// --------------------- END Refresh Token CRUD ------------------------ 

// --------------------- Create User ------------------------ 
func CreateUser(email string, password string, fullName string){
	tx := db.MustBegin()
    tx.MustExec(`
		INSERT INTO auth.users (email, password, full_name)
		VALUES ($1, $2, $3);
	`, email, password, fullName)
	
    tx.Commit()
} 
// --------------------- END Create User ------------------------ 


// --------------------- Verification Token CRUD ------------------------ 
type VerificationToken struct {
	ID        string `db:"id"`
	user_id   string
	TokenHash string `db:"token_hash"`
}

func GetEmailVerificationToken(email string)(string, error){
	tokenStr := VerificationToken{} 
	err := db.Get(&tokenStr,`
		SELECT * 
		FROM auth.verification_tokens
		WHERE email = "$1"
	`, email)
	if err != nil {
		if err == sql.ErrNoRows {
        	// Handle the "Not Found" case specifically
        	fmt.Println("No Verification code matched email passed")
        	return "", nil
    	}
		return "", err
	}
	return tokenStr.TokenHash, nil 
}

func SetUserToEmailVerified(userId string){
	tx := db.MustBegin()
    tx.MustExec(`
		UPDATE auth.users
		SET is_email_verified = TRUE
		WHERE user_id = $1;
	`, userId)
	
    tx.Commit()
}

func StoreVerificationToken(userId string, tokenStr string) {
	// calculating expires_at value
	const tokenExpiresIn = time.Minute * 15 
	expiresAt := time.Now().Add(tokenExpiresIn).Format(time.RFC3339)  
    tx := db.MustBegin()
    tx.MustExec(
		`
			INSERT INTO auth.verification_tokens 
			(user_id, token_hash, expires_at) VALUES ($1, $2, $3)
		`, 
		userId, tokenStr, expiresAt)
    tx.Commit()
}
// --------------------- END Verification Token CRUD ------------------------ 
