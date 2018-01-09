package main

import (
	"net/http"
)

func secureMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	appOrigin := r.URL.Query()["app-origin"]
	if len(appOrigin) < 1 || appOrigin[0] != "http://localhost:8081" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"error": "app unauthorized"}`))
		return
	}

	if _, err := r.Cookie("session"); err != nil || r.URL.Query()["username"][0] != "user1@example.com" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "user unauthorized"}`))
		return
	}

	w.Write([]byte(`{"id": "123", "message": "Secure message"}`))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "session", Value: "random value", MaxAge: 30, HttpOnly: true}
	http.SetCookie(w, &cookie)
	w.Write([]byte("cookie set for 30 seconds"))
}

func expireCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "session", MaxAge: -1}
	http.SetCookie(w, &cookie)
	w.Write([]byte("cookie expired"))
}

func main() {
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/securemessage", secureMessage)
	apiMux.HandleFunc("/setcookie", setCookie)
	apiMux.HandleFunc("/expirecookie", expireCookie)

	publicApi := http.StripPrefix("/", http.FileServer(http.Dir("public-api/")))
	apiMux.Handle("/", publicApi)

	go http.ListenAndServe(":8080", apiMux)

	publicApp := http.StripPrefix("/", http.FileServer(http.Dir("public-app/")))

	appMux := http.NewServeMux()
	appMux.Handle("/", publicApp)
	go http.ListenAndServe(":8081", appMux)
	http.ListenAndServe(":8082", appMux)
}
