package database



import (
	"path/filepath"
	//"encoding/binary"
	//"encoding/json"
	//"bufio"
	"log"
	//"fmt"
	"os"
	//"bytes"
)


var logPath string



func GetAbsPath(d* Database) (string, error) {

	logPath = "/mnt/c/Users/Dylan/My Documents/self_learning/DCache/database/" + d.Name + ".txt"

	absPath, err := filepath.Abs(logPath)
	if err != nil {
		return "", err
	}

	return absPath, nil

}



func UpdateLog(d *Database, cmd string) {

	// over writes the log file with the current state of the database address space

	// locate the log file first

	file, err := os.OpenFile(d.Name + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	// whenever a cmd is sent, it is logged in the file 

	file.WriteString(cmd + "\n")


	return
}



func LoadLog(d *Database) {

	// loads a log file to the address space

	// locate the log file first

	file, err := os.Open(d.Name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	return


}