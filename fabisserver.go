package main

import (
	"log"
	"net/http"
	"strings"
)

//Global declaration

var processName [2] string


//Programm
func mainHandler(w http.ResponseWriter, r *http.Request) {
q := r.URL.Query()
name := "Hell"
if len(q) == 0 {
name = "World"
}
responseString := "<html><body>Hello " + name + "</body></html>"
w.Write([]byte(responseString))
}

func startSite(w http.ResponseWriter, r *http.Request) {
//	url := r.URL.String()
	responseString := load()
	w.Write([]byte(responseString))
}

func load () string{
	//Laden der einzelen Prozesse .xml 
	processName[0] = "Notepad"
	processName[1] = "Calc"
	
	
	//Generieren der HTML Seite
	htmlPage := "<!DOCTYPE html><html><head><style type=\"text/css\">hr {"+
	"width: 95%;"+ 
    "height: 5px;"+
    "margin: 0 auto;"+
    "color: blue;"+
    "background: #dfac20;"+
    "}</style></head><body>" 
    
	for i := 0; i < 2; i++ {
		s1 := string (i+1)
		s2 := string (i+2)
		htmlPage = htmlPage + "<form action=\"/start/"+ processName[i]+"\">"+
 		"<button name="+ processName[i] + "type=\"submit\" value=\""+ processName[i]+"\">" + processName[i] + "</button>"+
		"<label for=\"check"+s1+"\"><input type=\"checkbox\" value=\"On\" id=\"check"+s1+"\">ON</label>"+
		"<label for=\"check"+s2+"\"><input type=\"checkbox\" value=\"Off\" id=\"check"+s2+"\">Off</label></form>"

		}
	htmlPage  = htmlPage + "</br><hr></br>"
	htmlPage = htmlPage + "Laufende Prozesse"
	//Einfügen laufender Prozesse
	htmlPage = htmlPage + "</body></html>"
	
	return htmlPage
}

func startProcess(w http.ResponseWriter, r *http.Request) {	
	url := strings.Trim(r.URL.Path, "/start/")			
	responseString := load()
	for i := 0; i < 2; i++ {
			if processName[i] == url {
				responseString = strings.Trim(responseString, "</body></html>")	
				responseString = responseString + "</br>" + processName[i]
			}
	}
	responseString = responseString + "</body></html>"
	w.Write([]byte(responseString))
	
}

func main() {
	http.HandleFunc("/", startSite)
	http.HandleFunc("/start/", startProcess)
//	log.Fatalln(http.ListenAndServeTLS(":1443", "cert.pem", "key.pem", nil))
	log.Fatalln(http.ListenAndServe(":8080",nil))
}
