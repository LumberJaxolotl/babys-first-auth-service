package lib

import (
	"log"
	"os"
	"time"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/fakedbhelpers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var JWT_SECRET string

func init() {
	// Loads the .env file and initialize JWT_SECRET as package-level variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWT_SECRET = os.Getenv("jwt_secret")
}

// return basic jwt string containing the user id
func GetUserAccessToken(userId string) string {

	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),                       // Issued at
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // Expires in 15 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenValue, err := token.SignedString(JWT_SECRET)
	if err != nil {
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
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // Expires in 15 minutes
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fakedbhelpers.StoreRefreshToken(token)

	tokenValue, err := token.SignedString(JWT_SECRET)

	if err != nil {
		log.Fatal("Error Signing Access Token")
	}

	return token, tokenValue, nil
}
