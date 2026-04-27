package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/dbhelpers"
	"github.com/LumberJaxolotl/babys-first-auth-service/models/fakedbhelpers"
	"github.com/LumberJaxolotl/babys-first-auth-service/models/lib"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	// quick check that all values are valid
	// err := r.ParseForm()
	// if err != nil {
	//     http.Error(w, "Failed to parse form", http.StatusBadRequest)
	//     return
	// }

	// name := r.Form.Get("name")
	// email := r.Form.Get("email")
	// password := r.Form.Get("password")

	// TODO fetch user id from verification token claims
	_, verificationTokenValue, _ := lib.GetUserVerificationToken("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")

	// fakedbhelpers.StoreVerificationToken(verificationTokenValue)

	cookie1 := &http.Cookie{
		Name:     "verification_token",
		Value:    verificationTokenValue,
		Path:     "/",  // For whole-site coverage
		HttpOnly: true, // Prevents JavaScript access (XSS protection)
		Secure:   true, // Ensures cookie is sent over HTTPS only
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600, // Expires in 1 hour (in seconds)
	}
	http.SetCookie(w, cookie1)

	// END ACCESS TOKEN COOKIE

	// learn more about process of email token verification
	// https://gemini.google.com/app/5135f941d0b938af

	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Email Verification</title>
</head>
<body>
    <form action="/verify-email" method="POST">
        <label for="code">Verification Code:</label>
        <input type="text" id="code" name="email_verification_code" required>
        <button type="submit">Submit</button>
    </form>
</body>
</html>`))
}

func VerifyEmailController(w http.ResponseWriter, r *http.Request) {

	// Check Code Passed is valid
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	email := r.Form.Get("email")
	emailVerificationCode := r.Form.Get("email_verification_code")

	dbRetrievedToken, err := dbhelpers.GetEmailVerificationToken(email)
	if err != nil {
		http.Error(w, "Failed to fetch token from database", 500)
		return
	}
	if dbRetrievedToken == "" {
		http.Error(w, "No matching token found", http.StatusNoContent)
		return
	}

	doTokensMatch, err := lib.DoTokensMatch(emailVerificationCode, dbRetrievedToken, sign)
	// TODO fetch user id from db after verifying stored verification token
	if err != nil {
		http.Error(w, "Error Verifying Email Verification Token", http.StatusBadRequest)
	}
	if doTokensMatch == false {
		http.Error(w, "Invalid Email Verification Token Passed", http.StatusUnauthorized)
	}

	accessTokenValue := lib.GetUserAccessToken("safdsafdsafdsafdsafdsafs")

	cookie1 := &http.Cookie{
		Name:     "access_token",
		Value:    accessTokenValue,
		Path:     "/",  // For whole-site coverage
		HttpOnly: true, // Prevents JavaScript access (XSS protection)
		Secure:   true, // Ensures cookie is sent over HTTPS only
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600, // Expires in 1 hour (in seconds)
	}
	http.SetCookie(w, cookie1)

	// END ACCESS TOKEN COOKIE

	// TODO store refresh token in db
	refreshToken, refreshTokenValue, _ := lib.GetUserRefreshToken("0")
	fakedbhelpers.StoreRefreshToken(refreshToken)

	cookie2 := &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshTokenValue,
		Path:     "/",  // For whole-site coverage
		HttpOnly: true, // Prevents JavaScript access (XSS protection)
		Secure:   true, // Ensures cookie is sent over HTTPS only
		SameSite: http.SameSiteLaxMode,
		MaxAge:   1209600, // Expires in 14 days (in seconds)
	}
	http.SetCookie(w, cookie2)

	w.Write([]byte("Hi, Logged in user"))
}

// For testing reference
var UNHASHED_PASSWORDS = []string{
	"P@ssw0rd123",
	"SecureKey!99",
	"QueryMaster",
	"DevOps_Life2026",
	"BlueSky$88",
	"ChartLover!22",
	"Nebula_77!#",
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	// TODO if both acces and refresh tokens, redirect to homepage

}
func RefreshTokenController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
func LogoutController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}

// TODO reimpliment getting jwt claims from header and then a
// db request, not through postgREST service
func GetMeController(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:3000/users?id=eq.a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Read the entire body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read body: %s\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
