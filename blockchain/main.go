package main

import (
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const blocksBucket = "blocks"
const dbFile = "./transactions/blockchain.db"

type Dbstatus struct {
	IsExists bool
	Status   os.FileInfo
	Path     string
}

func (d *Dbstatus) dbExists() *Dbstatus {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		d.IsExists = false
		d.Status, _ = os.Stat(dbFile)
		return d
	}
	d.IsExists = true
	d.Path = dbFile
	d.Status, _ = os.Stat(dbFile)
	return d
}

func main() {
	b := new(Dbstatus)
	r := b.dbExists()
	if r.IsExists == true {
		fmt.Println(r.IsExists)
	} else {
		fmt.Println("Can not find file.")
	}

	db, err := bolt.Open(b.Path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		fmt.Println(*b)
		return nil
	})
	defer db.Close()
}
