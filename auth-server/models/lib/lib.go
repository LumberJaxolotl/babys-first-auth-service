package lib

import (
	"crypto/sha256"
	"errors"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// TODO split file into JWTHELPERS and lib.go



var VERIFICATION_TOKEN_SIGNING_SECRET []byte
var ACCESS_TOKEN_SIGNING_SECRET []byte
var REFRESH_TOKEN_SIGNING_SECRET []byte

// idomatic way to declare enums
type TokenType string
const (
	TokenTypeVerification TokenType = "email_verification"
	TokenTypeRefresh      TokenType = "refresh"
	TokenTypeAccess       TokenType = "access"
	Token
)

var TOKEN_SIGNING_SECRETS map[TokenType][]byte




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
	} else if accessTokenSigningSecret == "" {
		log.Fatal("Enviroment variable 'ACCESS_TOKEN_SIGNING_SECRET' could not be found")
	} else if refreshTokenSigningSecret == "" {
		log.Fatal("Enviroment variable 'REFRESH_TOKEN_SIGNING_SECRET' could not be found")
	}

	VERIFICATION_TOKEN_SIGNING_SECRET = []byte(verificationTokenSigningSecret)
	ACCESS_TOKEN_SIGNING_SECRET = []byte(accessTokenSigningSecret)
	REFRESH_TOKEN_SIGNING_SECRET = []byte(refreshTokenSigningSecret)

	// for convenience later in testing :)
	TOKEN_SIGNING_SECRETS = map[TokenType][]byte{
		"verification": VERIFICATION_TOKEN_SIGNING_SECRET,
		"access": ACCESS_TOKEN_SIGNING_SECRET,
		"refresh": REFRESH_TOKEN_SIGNING_SECRET,
	}
}

// ---- Verification Tokens | Generation and Verification Logic -----
func GenerateUserVerificationToken(userId string) (*jwt.Token, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(),                       // Issued at
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // times-out at
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims), nil
}

// ---- END Access Tokens | Generation and Verification Logic -----

// ---- Access Tokens | Generation and Verification Logic -----

// returns basic jwt string containing the user id
func GenerateUserAccessToken(userId string) *jwt.Token {
	// TODO impliment
	return nil	
}

// ---- Refresh Tokens | Generation and Verification Logic -----

// Returns refresh token obj for storing in DB and
// string value for use in cookie
func GenerateUserRefreshToken(userId string) *jwt.Token {
	// TODO impliment
	return nil
}

// ------------------------- Misc. Helpers ------------------------------

func HashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	bcryptHash, err := bcrypt.GenerateFromPassword(hash[:], bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptHash), nil
}



func GenerateAuthToken(userId string, tokenType TokenType, validFor time.Duration) (*jwt.Token, string, error) {

	claims := jwt.MapClaims{
		"user_id": userId,
		"iat":     time.Now().Unix(), // Issued at
		"exp":     time.Now().Add(time.Minute * validFor).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signingSecret := TOKEN_SIGNING_SECRETS[tokenType]

	signedTokenHash, err := token.SignedString(signingSecret)
	if err != nil {
		return nil, "", errors.New("Error Signing " + string(tokenType) + " Token")
	}

	return token, signedTokenHash, nil
}

func DoTokensMatch(token1 string, token2 string, signingSecret string) (bool, error) {
	// TODO I think hashes need to be salted and hashed again before being put in the database
	return token1 == token2, nil
}
