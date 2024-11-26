package database

import (
	//"fmt"
	"bytes"
	"time"
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
	Type        EntryType
	Key         string
	Data        *bytes.Buffer
	Next        *Entry         // entries may be chained together if there is a collision
	UpdatedAt   time.Time
}



type Database struct {
	AddressSpace []Entry
	MaxSize      uint32
}



func (d *Database) InitDatabase(size uint32) {
	// size is num entries
	d.AddressSpace = make([]Entry, size)  // current size of 0, capacity of size
	d.MaxSize = size

	return 
}


// define a hash function to map keys to indexes
// index = sum(ASCII values in key) modulo 599


func (d *Database) Hash(key string) int {

	s := 0

	for _, char := range key {

		s += int(char)

	}

	return s % int(d.MaxSize)

}



// Add functionality to add Entries to the database
func (d *Database) AddEntry(entryType EntryType, key string, data *bytes.Buffer) error {


	// make a new entry
	// going to need to check if key is already in the database


	present, presentEntry := d.GetEntry(key)

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
 	

 	index := d.Hash(key) 

 	d.AddressSpace[index] = entry


 	return nil
}



func (d *Database) GetEntry(key string) (bool, *Entry) {

	// returns true, Entry or false Empty Entry

	index := d.Hash(key) // this is the index that the entry will be at if it is there

	e := d.AddressSpace[index] // entry containing potential list of entries



	for e.Key != key && e.Next != nil {
		e = *e.Next
	}

	if e.Key == "" {
		// not in the database
		emptyEntry := Entry{}
		return false, &emptyEntry
	}

	return true, &e

}
