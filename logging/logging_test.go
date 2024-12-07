package logging


import (
	"bytes"
	"testing"
	"github.com/DylanCoon99/DCache/database"
)



func TestUpdateLog(t *testing.T) {

	var size uint32

	size = 10
	name := "database1"

	var d *database.Database

	d = database.InitDatabase(name, size)


	d1 := bytes.NewBuffer([]byte("This is test data for key1"))

	database.AddEntry(database.StringType, "key1", d1)


	d2 := bytes.NewBuffer([]byte("This is test data for key2"))

	database.AddEntry(database.StringType, "key2", d2)



	UpdateLog(d)

	LoadLog(d)

}