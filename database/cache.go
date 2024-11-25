package database

import (
	"fmt"
	"bytes"
)


// This is where all of the database structs will be defined

/* Think about the design of the database

	Key Value Store:

		Hash Map Cache
			- 





*/


type EntryType int


const (
	StringType EntryType = iota
	ListType   
	SetType
	HashType
	SortedSetType
	BitmapType
)


type Entry struct {
	Type      EntryType
	Key       string
	Data      *bytes.Buffer
}



type Database struct {
	AddressSpace []Entry
	MaxSize      uint32
}



func (d *Database) InitDatabase(size uint32) {
	// size is num entries
	d.AddressSpace = make([]Entry, size)  // current size of 0, capacity of size
	d.MaxSize = size

	fmt.Println(d.AddressSpace)

	return 
}


// define a hash function to map keys to indexes
// index = sum(ASCII values in key) modulo 599


func Hash(key string) int {

	s := 0

	for _, char := range key {

		s += int(char)

	}

	return s % 599

}



// Add functionality to add Entries to the database
func (d *Database) AddEntry(entryType EntryType, key string, data *bytes.Buffer) error {


	// make a new entry
	/*
	entry := Entry {
		Type: entryType,
		Key: key,
		Data: data,
 	}
 	*/

 	//fmt.Println(entry.Type)

 	//i := 0

 	//d.AddressSpace[i] = entry


 	return nil
}