package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Timestamp records the currency time when the block be created.
// Data is the useful information in this Block
// PrevBlockHash stored the hash of previous Block
// Hash

// These for variable above is called block headers; Although Transactions,
// which called Data in this scripy, mostly constructed by another data struct
// but I build this Block strcut briefly.

// Then we may think, how to caculate hash code? Since hash code is essnetial to block and
// of course, Getting a hash code is a computational complexity problem--even thoutht
// the multi-core high preformance computer have to take lots of time(That's why are there some people buy more GPU to improve the computational ability)

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// The function above is use to connect different value between different block.

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// Create a block

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func init() {
	fmt.Println("initialization...")
}
