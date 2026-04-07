package main

import (
	//...

	"fmt"
	"net/http"
	"time"

	"github.com/LumberJaxolotl/babys-first-auth-service/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	})

	r.Route("/auth", func(r chi.Router) {
		// 1. Public: Register a new user
		r.Post("/register", controllers.RegisterController)

		// 3. Public: Use Refresh Token to get a new Access Token
		r.Post("/refresh", controllers.RefreshTokenController)
		
		// 2. Public: Exchange credentials for tokens
		r.Post("/login", controllers.LoginController)

		// 4. Public/Private: Invalidate the session
		r.Post("/logout", controllers.LogoutController)

		// 5. Protected: Get current user info (requires Auth middleware)
		r.Get("/me", controllers.GetMeController)

	})

	fmt.Println("Auth server running on port 80")
	http.ListenAndServe(":80", r)
}
