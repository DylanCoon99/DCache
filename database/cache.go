package database

import (
	"bytes"
	"time"
	//"fmt"
)


// This is where all of the database structs will be defined

/* Think about the design of the database

	Key Value Store:

		Hash Map Cache
			- 



12/2 to do
 - figure out how to get data to persist (done)


12/7 to do
 - when initing a database -> read the log.txt file (done)
 - fix the GetEntry functionality (done)
 - implement different types
 	- think about how to parse types

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
	Type        EntryType     `json:"type"`
	Key         string        `json:"key"`
	Data        *bytes.Buffer `json:"data"`
	Next        *Entry        `json:"next"`       // entries may be chained together if there is a collision
	UpdatedAt   time.Time     `json:"updated_at"`
}



type Database struct {
	Name         string
	AddressSpace []Entry
	MaxSize      uint32
}


var D *Database


func InitDatabase(name string, size uint32) *Database {
	// size is num entries
	d := Database{}



	d.AddressSpace = make([]Entry, size)  // current size of 0, capacity of size
	d.MaxSize = size
	d.Name = name

	D = &d

	return D
}


// define a hash function to map keys to indexes
// index = sum(ASCII values in key) modulo 599


func Hash(key string) int {

	s := 0

	for _, char := range key {

		s += int(char)

	}

	return s % int(D.MaxSize)

}



// Add functionality to add Entries to the database
func AddEntry(entryType EntryType, key string, data *bytes.Buffer) error {


	// make a new entry
	// going to need to check if key is already in the database

	defer UpdateLog(D)


	present, presentEntry := GetEntry(key)


	if present {

		presentEntry.Type = entryType
		presentEntry.Data = data
		presentEntry.UpdatedAt = time.Now()


		return nil
	}
	
	entry := Entry {
		Type: entryType,
		Key: key,
		Data: data,
		UpdatedAt: time.Now(),
 	}
 	

 	index := Hash(key) 


 	// need to check if a different entry already exists at this index


 	loc := &D.AddressSpace[index]

 	if loc.Key == "" {
 		// if this space is empty
 		D.AddressSpace[index] = entry

 		return nil
 	}


 	for loc.Next != nil {
 		loc = loc.Next
 	}

 	loc.Next = &entry


 	return nil
}



func GetEntry(key string) (bool, *Entry) {

	// returns true, Entry or false Empty Entry

	index := Hash(key) // this is the index that the entry will be at if it is there

	e := &D.AddressSpace[index] // entry containing potential list of entries



	for e.Key != key && e.Next != nil {
		e = e.Next
	}

	if e.Key == "" || e.Key != key {
		// not in the database
		emptyEntry := Entry{}
		return false, &emptyEntry
	}


	return true, e

}


