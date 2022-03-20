package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "firstName",
		Value:    "firstValue",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "secondName",
		Value:    "secondValue",
		HttpOnly: true,
	}

	w.Header().Set("Set-Cookie", c1.String())

	// use Add not Set, otherwise will overwrite c1
	w.Header().Add("Set-Cookie", c2.String())
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
