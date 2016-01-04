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
func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/test", test)
	log.Fatalln(http.ListenAndServeTLS(":1443", "cert.pem", "key.pem", nil))
}
