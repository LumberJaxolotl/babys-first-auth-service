package controllers

import (
	"net/http"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
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