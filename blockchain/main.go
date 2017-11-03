package main

import (
	"fmt"
	"os"

	_ "github.com/boltdb/bolt"
)

const dbFile = "./transactions/blockchain.db"

type Dbstatus struct {
	IsExists bool
	Status   os.FileInfo
}

func (d *Dbstatus) dbExists() *Dbstatus {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		d.IsExists = false
		d.Status, _ = os.Stat(dbFile)
		return d
	}
	d.IsExists = true
	d.Status, _ = os.Stat(dbFile)
	return d
}

func main() {
	b := new(Dbstatus)
	r := b.dbExists()
	fmt.Println(*r)
}
