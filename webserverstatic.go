package main

import (
	"log"
	"net/http"
)
func mainHandler(w http.ResponseWriter, r *http.Request) {
q := r.URL.Query()
name := "Hell"
if len(q) == 0 {
name = "World"
}
responseString := "<html><body>Hello " + name + "</body></html>"
w.Write([]byte(responseString))
}

func mainHand(w http.ResponseWriter, r *http.Request) {
//	url := r.URL.String()
	function := "<form action=/test><input type='submit' value='Go to Google'></form>"
	responseString := "<html><body><button>Hello " + function + "</button></body></html>"
	w.Write([]byte(responseString))
}

func test(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><h1>Hello World</body></html>"
	w.Write([]byte(responseString))
}

func peterfunc(w http.ResponseWriter, r *http.Request) {
	responseString := "<html><h1>Hello World</body></html>"
	w.Write([]byte(responseString))
}

func main() {
	//http.HandleFunc("/", mainHandler)
	//http.HandleFunc("/test", test)
	//http.HandleFunc("/peter", peterfunc)
	//http.Handle("/peter", http.FileServer(http.Dir("html")))
	
	  fs := http.FileServer(http.Dir("html"))
http.Handle("/peterle/", http.StripPrefix("/peterle/", fs))
	//http.Handle("/peter",    http.ServeFile("/html/gowut.html"))
	log.Fatalln(http.ListenAndServe(":3000", nil))
	//log.Fatalln(http.ListenAndServeTLS(":1443", "cert.pem", "key.pem", nil))
}
