package controllers

import (
	"net/http"

	"github.com/LumberJaxolotl/babys-first-auth-service/models/lib"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {


	tokenValue := lib.GetUserAccessToken("0")

	cookie1 := &http.Cookie{
        Name:     "access_token",
        Value:    tokenValue,
        Path:     "/",             // For whole-site coverage
        HttpOnly: true,            // Prevents JavaScript access (XSS protection)
        Secure:   true,            // Ensures cookie is sent over HTTPS only
        SameSite: http.SameSiteLaxMode, 
        MaxAge:   3600,            // Expires in 1 hour (in seconds)
    }
    http.SetCookie(w, cookie1)

	// END ACCESS TOKEN COOKIE 



	cookie2 := &http.Cookie{
        Name:     "access_token",
        Value:    tokenValue,
        Path:     "/auth/refresh", // For whole-site coverage
        HttpOnly: true,            // Prevents JavaScript access (XSS protection)
        Secure:   true,            // Ensures cookie is sent over HTTPS only
        SameSite: http.SameSiteLaxMode, 
		MaxAge:   1209600,            // Expires in 14 days (in seconds)
    }
    http.SetCookie(w, cookie2)	
	
	w.Write([]byte("Hi, Logged in user"))
}
func LoginController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
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