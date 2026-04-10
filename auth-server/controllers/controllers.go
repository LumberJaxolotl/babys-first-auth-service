package controllers

import (
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	


	accessToken := "your-secure-value"
	cookie := &http.Cookie{
        Name:     "access_token",
        Value:    accessToken,
        Path:     "/",             // For whole-site coverage
        HttpOnly: true,            // Prevents JavaScript access (XSS protection)
        Secure:   true,            // Ensures cookie is sent over HTTPS only
        SameSite: http.SameSiteLaxMode, 
        MaxAge:   3600,            // Expires in 1 hour (in seconds)
    }
    http.SetCookie(w, cookie)	
	
	
	
	w.Write([]byte("Hi. "))
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