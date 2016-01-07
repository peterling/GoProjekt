package main

import (
    "fmt"
    "io"
    "log"
    "os"
    "os/exec"
	//"os/signal"
)

func checkError(err error) {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
} 
//zuerst aus XML die Startbefehle einlesen in Applikation-struct
type Applikation struct {
    startCommand string
    startArgument string
	startAttribut string
	stopCommand string
	stopArgument string
	stopAttribut string
	}

func main() {
	//Liste aller verfügbaren bzw. startbaren Applikationen
	 aplliste := [3]Applikation{
        {startCommand: "ping", startArgument: "8.8.8.8", startAttribut: "-4",
		stopCommand: "ping", stopArgument:"8.8.4.4", stopAttribut:"-6"},
       {startCommand: "ping", startArgument: "localhost", startAttribut: "-4",
		stopCommand: "ping", stopArgument:"8.8.4.4", stopAttribut:"-6"},
      {startCommand: "ping", startArgument: "www.google.de", startAttribut: "-4",
		stopCommand: "ping", stopArgument:"8.8.4.4", stopAttribut:"-6"},  }

	//Irgendwie für jede Applikation einen exec.Command erstellen... nur wie ?
	for index,element := range aplliste {
	//	cmd[index]:=exec.Command(aplliste[index].startCommand, aplliste[index].startArgument)
		fmt.Println()
		fmt.Print(index)
		fmt.Print(element)
		fmt.Println()
}	//Hier sind die exec.Commands manuell eingerichtet, diese werden später aufgerufen
    cmd := exec.Command(aplliste[0].startCommand, aplliste[0].startArgument,aplliste[0].startAttribut)
	cmd_peter:=exec.Command(aplliste[1].startCommand, aplliste[1].startArgument,aplliste[1].startAttribut)
    //cmd_peter := exec.Command("cmd")
	// Create stdout, stderr streams of type io.Reader
	
    //Standard-Ausgabe und -Fehler umleiten - momentan noch auf die Konsole...
	stdout, err := cmd.StdoutPipe()
    checkError(err)
    stderr, err := cmd.StderrPipe()
    checkError(err)
	stdout_peter, err_peter := cmd_peter.StdoutPipe()
    checkError(err_peter)
	
	/*	stdin, err := cmd_peter.StdinPipe()
if err != nil {
    fmt.Println(err)
}
defer stdin.Close()*/
//io.WriteString(stdin, "dir\n")
//io.WriteString(stdin, "dir\n")
//io.WriteString(stdin, "exit\n")
//signal.Notify(stdin, os.Interrupt)
	
	
	//Standard-Eingabe umleiten. Evtl. garnicht notwendig, wenn wir mit SIGINT arbeiten können...
	stdin_peter, err_peter := cmd_peter.StdinPipe()
	checkError((err_peter))
    stderr_peter, err_peter := cmd_peter.StderrPipe()
    checkError(err_peter)

    //Die oben angelegten exec.Commands ausführen.
    err = cmd.Start()
    checkError(err)
	err_peter = cmd_peter.Start()
	checkError(err_peter)

    // main() darf erst beendet werden, wenn unsere Commands fertig sind...
    defer cmd.Wait()  // blockt nicht, läuft also weiter
	defer cmd_peter.Wait()
	
    // nicht-blockende echo-Ausgabe, momentan zum Terminal... (go ...)
    go io.Copy(os.Stdout, stdout)
    go io.Copy(os.Stderr, stderr)
    go io.Copy(os.Stdout, stdout_peter)
	go io.Copy(stdin_peter, os.Stdin)
	//go io.Copy(ioutil.ReadAll(os.Stdin), stdin_peter)

//Auslesen der PIDs zum Steuern der Prozesse... muss noch erforscht werden
//fmt.Println(cmd_peter.Process.Pid)
fmt.Println(cmd_peter.Process.Pid)
fmt.Println(cmd.Process.Pid)
fmt.Println(os.Getpid()) //Achtung, PID der Hauptanwendung!

//unter Windows funktioniert nur Kill, nicht Interrupt.
cmd.Process.Signal(os.Interrupt)
cmd_peter.Process.Signal(os.Interrupt)
//cmd.Process.Signal(os.Kill)
//cmd_peter.Process.Signal(os.Kill)

   go io.Copy(os.Stderr, stderr_peter)
    // I love Go's trivial concurrency :-D
    fmt.Printf("Hier kann jetzt parallel anderer Kram passieren.\n\n")
}
//p,_:=os.FindProcess(cmd_peter.Process.Pid)
//p.Kill()
//fmt.Println(cmd_peter.ProcessState.Pid())
	//go io.PipeWriter(os.Stdin, stdin_peter)
 
