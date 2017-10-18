package auth

import "net/http"

// LogoutHandler is delete auth cookie
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "auth",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	w.Header().Set("Location", "/default")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
