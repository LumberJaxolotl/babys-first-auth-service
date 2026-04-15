package controllers

import (
	"net/http"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/fakedbhelpers"
	"github.com/LumberJaxolotl/babys-first-auth-service/models/lib"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	
	
	// TODO fetch user id from verification token claims
	_, verificationTokenValue, _ := lib.GetUserVerificationToken("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")
	
	cookie1 := &http.Cookie{
        Name:     "verification_token",
        Value:    verificationTokenValue,
        Path:     "/",             // For whole-site coverage
        HttpOnly: true,            // Prevents JavaScript access (XSS protection)
        Secure:   true,            // Ensures cookie is sent over HTTPS only
        SameSite: http.SameSiteLaxMode, 
        MaxAge:   3600,            // Expires in 1 hour (in seconds)
    }
    http.SetCookie(w, cookie1)

	// END ACCESS TOKEN COOKIE 

	fakedbhelpers.StoreVerificationToken(verificationTokenValue)


	w.Write([]byte("Hi, Logged in user"))
}

func VerifyEmailController(w http.ResponseWriter, r *http.Request) {

	// TODO fetch user id from db after verifying stored verification token  
	accessTokenValue := lib.GetUserAccessToken()
	
	cookie1 := &http.Cookie{
        Name:     "access_token",
        Value:    accessTokenValue,
        Path:     "/",             // For whole-site coverage
        HttpOnly: true,            // Prevents JavaScript access (XSS protection)
        Secure:   true,            // Ensures cookie is sent over HTTPS only
        SameSite: http.SameSiteLaxMode, 
        MaxAge:   3600,            // Expires in 1 hour (in seconds)
    }
    http.SetCookie(w, cookie1)

	// END ACCESS TOKEN COOKIE 

	// TODO store refresh token in db
	refreshToken, refreshTokenValue, _ := lib.GetUserRefreshToken("0")
	fakedbhelpers.StoreRefreshToken(refreshToken)

	cookie2 := &http.Cookie{
        Name:     "refresh_token",
        Value:    refreshTokenValue,
        Path:     "/", // For whole-site coverage
        HttpOnly: true,            // Prevents JavaScript access (XSS protection)
        Secure:   true,            // Ensures cookie is sent over HTTPS only
        SameSite: http.SameSiteLaxMode, 
		MaxAge:   1209600,            // Expires in 14 days (in seconds)
    }
    http.SetCookie(w, cookie2)	
	
	w.Write([]byte("Hi, Logged in user"))
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
func GetMeController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}