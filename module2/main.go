package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	for key, value := range r.Header {
		w.Header().Set(key, value[0])
	}
	os.Setenv("VERSION", "0.0.1")
	env := os.Getenv("VERSION")
	fmt.Fprint(w, env)
	log.Printf("%s %s", r.RemoteAddr, r.Method)
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", Healthz)
	http.ListenAndServe(":8000", nil)
}
