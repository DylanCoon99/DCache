package logging

// logging package is responsible for writing/loading to/from log file


import (
	"github.com/DylanCoon99/DCache/database"
	"encoding/json"
	"bufio"
	"log"
	"fmt"
	"os"
)


func UpdateLog(d *database.Database) {

	// over writes the log file with the current state of the database address space

	// locate the log file first

	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// iterate over the address space
	addressSpace := d.AddressSpace

	for i := 0; i < len(addressSpace); i ++ {
		data, _ := json.Marshal(addressSpace[i])
		_, err = file.WriteString(string(data) + "\n")
		if err != nil {
			log.Fatal(err)
		}

	}

}



func LoadLog(d *database.Database) {

	// loads a log file to the address space

	// locate the log file first

	file, err := os.Open(d.Name + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}


}