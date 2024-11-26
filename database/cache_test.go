package database

import (
	"fmt"
	"testing"
	"bytes"
)


/*
func TestInitDatabase(t *testing.T) {


	var size uint32

	size = 10

	var d Database

	d.InitDatabase(size)


	fmt.Println(d.MaxSize)


}
*/



/*
func TestHash(t *testing.T) {

	ret := Hash("This is a bunch of nonsense")
	fmt.Printf("This is the result %d: ", ret)
}
*/





func TestAddEntry(t *testing.T) {

	var size uint32

	size = 10

	var d Database

	d.InitDatabase(size)

	d1 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key1", d1)

	d2 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key2", d2)

	d3 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key3", d3)

	d4 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key4", d4)

	d5 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key5", d5)

	d6 := bytes.NewBuffer([]byte("This is test data"))

	d.AddEntry(StringType, "key6", d6)

	for i := 0; i < len(d.AddressSpace); i ++ {
		fmt.Println(d.AddressSpace[i])
	}


	d6 = bytes.NewBuffer([]byte("This is some updated data for key6"))
	d.AddEntry(StringType, "key6", d6)
	fmt.Println("///////////////")

	for i := 0; i < len(d.AddressSpace); i ++ {
		fmt.Println(d.AddressSpace[i])
	}

}
