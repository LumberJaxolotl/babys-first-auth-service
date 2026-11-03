package dbhelpers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/lib"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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



// --------------------- Create User ------------------------ 
func CreateUser(email string, password string, fullName string)(string, error){
	// Generates a uuid
	uuid, err := uuid.NewV7()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	userId := uuid.String() 
	// Adds user to database
	tx := db.MustBegin()
    tx.MustExec(`
		INSERT INTO auth.users (id, email, password, full_name)
		VALUES ($1, $2, $3, $4);
	`, userId, email, password, fullName)
    tx.Commit()
	// Returns id of created user
	return userId, nil
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
	// part of quuery to put back in later
	// ON auth.verification_tokens.user_id = auth.users.id
	err := db.Get(&tokenStr,`SELECT token_hash 
		FROM auth.verification_tokens
		LEFT JOIN auth.users
			ON auth.users.id = auth.verification_tokens.user_id
		WHERE auth.users.email = 'alex.rivera@example.com';`)
	
	if err != nil {
		if err == sql.ErrNoRows {
        	// Handle the "Not Found" case specifically
        	fmt.Println("No Verification Token matched email passed")
        	return "", nil
    	}
		return "", err
	}
	return tokenStr.TokenHash, nil 
}

// sets user as verified, 
// deletes associated user verification token and verification code, 
func SetUserToEmailVerified(userId string){
	tx := db.MustBegin()
    tx.MustExec(`
		UPDATE auth.users
		SET is_email_verified = TRUE
		WHERE user_id = $1;
	`, userId)
	// todo add sql command to delete 
    tx.Commit()
}

func StoreVerificationToken(userId string, verificationToken *jwt.Token) {
	// calculating expires_at value
	tx := db.MustBegin()
    // takes the time to clear out any expired tokens so they  
	// dont fill the db
	tx.MustExec(`
    	DELETE FROM auth.verification_tokens 
    	WHERE user_id = $1 AND expires_at < $2
	`, userId, time.Now())
	

	signingSecret := lib.TOKEN_SIGNING_SECRETS["verification"]
	tokenHash, err := verificationToken.SignedString(signingSecret)
	if err != nil {
		log.Fatal("Error Signing Verification Token")
	}

	expiresAt, err :=  verificationToken.Claims.GetExpirationTime()

	tx.MustExec(`
			INSERT INTO auth.verification_tokens 
				(user_id, token_hash, expires_at) 
			VALUES ($1, $2, $3)
	`, userId, tokenHash, expiresAt.Time)
    tx.Commit()
}


// --------------------- END Verification Token CRUD ------------------------ 


// --------------------- Verification Code CRUD ------------------------ 


var emailVerificationCode int
// TODO Impliment
func IsCorrectEmailVerificationCode(email string, code string) (bool, error){
	codeFromUser, convertionErr := strconv.Atoi(code)
	if convertionErr != nil {
		return false, convertionErr
	}

	err := db.Get(&emailVerificationCode,`
		SELECT evc.email_verification_code  
		FROM auth.users u
		LEFT JOIN auth.email_verification_codes evc
			ON u.id = evc.user_id 
		WHERE u.email = '$1';
	`);
	if err != nil {
		if err == sql.ErrNoRows {
        	// Handle the "Not Found" case specifically
        	fmt.Println("No Verification Token matched email passed")
        	return false, nil
    	}
		return false, err
	}

	if codeFromUser != emailVerificationCode{
        return false, errors.New("No Verification Token matched email passed")
	}
	
	return true, nil
}

func StoreVerificationCode() error {
	// Todo impliment	
	return nil
}

// --------------------- END Verification Code CRUD ------------------------ 

// --------------------- Refresh Token CRUD ------------------------ 

func GenerateRefreshToken(tokenStr string) (*jwt.Token, error) {
	return nil, nil
}

func StoreRefreshToken(userId string, token *jwt.Token) {
	
}

// --------------------- END Refresh Token CRUD ------------------------ 