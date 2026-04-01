package handlers

import (
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}
func GetMeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi"))
}