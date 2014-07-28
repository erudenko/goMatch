package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":8080", "http server address")
var homeTemplate = template.Must(template.ParseFiles("Home.html"))
var userTemplate = template.Must(template.ParseFiles("user1.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}

func serveUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Path[len("/user/"):]

	if len(userID) == 0 {
		http.Error(w, "Not Found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	filter := UserFilter{UserID: userID}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	userTemplate.Execute(w, filter)
}

func main() {
	flag.Parse()
	go matcher.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/user/", serveUser)
	http.HandleFunc("/ws", serveSocket)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
