package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var VERIFICATION_TOKEN_SIGNING_SECRET []byte
var ACCESS_TOKEN_SIGNING_SECRET []byte
var REFRESH_TOKEN_SIGNING_SECRET []byte


func init() {
	// Loads the .env file and initializes the signing secrets as package-level variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	verificationTokenSigningSecret := os.Getenv("VERIFICATION_TOKEN_SIGNING_SECRET")
	accessTokenSigningSecret := os.Getenv("ACCESS_TOKEN_SIGNING_SECRET")
	refreshTokenSigningSecret := os.Getenv("REFRESH_TOKEN_SIGNING_SECRET")

	if verificationTokenSigningSecret == "" {
		log.Fatal("Enviroment variable 'VERIFICATION_TOKEN_SIGNING_SECRET' could not be found")
	}else if accessTokenSigningSecret == "" {
		log.Fatal("Enviroment variable 'ACCESS_TOKEN_SIGNING_SECRET' could not be found")
	}else if refreshTokenSigningSecret == "" {
		log.Fatal("Enviroment variable 'REFRESH_TOKEN_SIGNING_SECRET' could not be found")
	}

	VERIFICATION_TOKEN_SIGNING_SECRET = []byte(verificationTokenSigningSecret)
	ACCESS_TOKEN_SIGNING_SECRET = []byte(accessTokenSigningSecret)
	REFRESH_TOKEN_SIGNING_SECRET = []byte(refreshTokenSigningSecret)
	
}

func GetUserVerificationToken(userId string) (*jwt.Token, string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),                       // Issued at
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // times-out at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err := token.SignedString(VERIFICATION_TOKEN_SIGNING_SECRET)


	if err != nil {
		fmt.Println(err)
		log.Fatal("Error Signing Access Token")
	}

	return token, tokenValue, nil
}


// return basic jwt string containing the user id
func GetUserAccessToken(userId string) string {

	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),                       // Issued at
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // Expires in 15 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err := token.SignedString(ACCESS_TOKEN_SIGNING_SECRET)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error Signing Access Token")
	}

	return tokenValue
}

// Returns refresh token obj for storing in DB and
// string value for use in cookie
func GetUserRefreshToken(userId string) (*jwt.Token, string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),                       // Issued at
		"exp":     time.Now().Add(time.Hour * 24 * 60).Unix(), // Expires in 60 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err := token.SignedString(REFRESH_TOKEN_SIGNING_SECRET)


	if err != nil {
		fmt.Println(err)
		log.Fatal("Error Signing Access Token")
	}

	return token, tokenValue, nil
}


