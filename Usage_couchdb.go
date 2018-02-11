package main

import (
	"log"
	"os"
	"os/exec"
)

const DefaultDatabase = "http://localhost:5984"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Test() {
	output := exec.Command("curl", "-X", "GET", "https://example.com/")
	output.Stdout = os.Stdout
	output.Stderr = os.Stderr
	err := output.Run()
	check(err)
}

/*
func NewDatabase(name string) string {
	//os.Stdin.Write("curl -X PUT " + DefaultDatabase + "/" + name)
	output := exec.Command("curl", "-X", "PUT", DefaultDatabase+"/"+name)
	return output
}

func GetDatabase(name string) string {
	output := exec.Command("curl", "-X", "GET", DefaultDatabase+"/"+name)
	return output
}


/*func Create(database_name string, id int, jsonInput ) {
	//os.Stdin.Write("curl -X PUT " + DefaultDatabase + "/" + name)
	output := exec.Command("curl", "-X", "PUT", DefaultDatabase+"/"+database_name, "/"+id, "-d",
		"'" + jsonInput + "'")
	return output
}


func Read(name string) {
}

func Update(name string) {

}

func Delete(name string) {

}
*/
func main() {
	Test()
}

/*func main() {
	newDBCmd := flag.NewFlagSet("newdatabase", flag.ExitOnError)
	insertCmd := flag.NewFlagSet("insert", flag.ExitOnError)
	readCmd := flag.NewFlagSet("read", flag.ExitOnError)
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	delCmd := flag.NewFlagSet("del", flag.ExitOnError)

	switch os.Args[1] {
	case "newdatabase":
		err := newDBCmd.Parse(os.Args[2:])
		check(err)
	case "insert":
		err := insertCmd.Parse(os.Args[2:])
		check(err)
	case "read":
		err := readCmd.Parse(os.Args[2:])
		check(err)
	case "update":
		err := updateCmd.Parse(os.Args[2:])
		check(err)
	case "delete":
		err := delCmd.Parse(os.Args[2:])
		check(err)
	}

	if newDBCmd.Parsed() {

	}

	if insertCmd.Parsed() {

	}

	if readCmd.Parsed() {

	}

	if updateCmd.Parsed() {

	}

	if delCmd.Parsed() {

	}
}*/
