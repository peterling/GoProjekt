package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
	"strings"
)

func Home(w http.ResponseWriter, req *http.Request) {
    render(w, "index.html")
}

func About(w http.ResponseWriter, req *http.Request) {
    render(w, "about.html")
}

func Peter(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()  // parse arguments, you have to call this by yourself
    fmt.Println(r.Form)  // print form information in server side
    //fmt.Println("path", r.URL.Path)
    //fmt.Println("scheme", r.URL.Scheme)
   // fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
	
	welchesProgramm := strings.Join(r.Form["program"],"")
	wasTun:=r.Form["aktion"]
	fmt.Fprintln(w, welchesProgramm)
	fmt.Fprintln(w, wasTun)
	
	ProgrammListe := map[string]bool {
    "15": true,
    "42": true,
}
if ProgrammListe[welchesProgramm] {
    fmt.Fprintln(w, "Programm vorhanden, Befehl wird ausgeführt")
	fmt.Fprintln(w, "Programm "+ welchesProgramm+" wurde")
	 if strings.Join(wasTun,"")=="starten"{
        fmt.Fprintln(w, "gestartet")
    } else if strings.Join(wasTun, "")=="beenden"{
        fmt.Fprintln(w, "beendet")
    } else {
       fmt.Fprintln(w, "nix wurde gemacht, falscher Aufruf!")}
}else{fmt.Fprintln(w, "angegebenes Programm existiert nicht, Befehl nicht ausführbar!")}

    //fmt.Fprintf(w, "\n\nHello Fabian!\n") // send data to client side
	
}

func render(w http.ResponseWriter, tmpl string) {
    tmpl = fmt.Sprintf("templates/%s", tmpl)
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        log.Print("template parsing error: ", err)
    }
    err = t.Execute(w, "")
    if err != nil {
        log.Print("template executing error: ", err)
    }
}

func main() {
    http.HandleFunc("/", Home)
    http.HandleFunc("/about/", About)
	http.HandleFunc("/peter", Peter)
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
