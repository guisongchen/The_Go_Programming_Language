package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	num := rand.Intn(10)
	fmt.Printf("num:%d\n", num)
	t.Execute(w, num > 5)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
