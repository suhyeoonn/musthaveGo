package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"a", 16, 87}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.HandleFunc("/bar", barHandler)
	mux.HandleFunc("/student", StudentHandler)

	return mux
}

func main() {
	//http.ListenAndServe(":3000", MakeWebHandler())
	err := http.ListenAndServeTLS(":3000", "server.crt", "server.key", MakeWebHandler())

	if err != nil {
		log.Fatal(err)
	}
	//http.Handle("/", http.FileServer(http.Dir("static")))
	//http.ListenAndServe(":3000", nil)
}
