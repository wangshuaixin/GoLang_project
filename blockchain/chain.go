package main

// now we have a block and know how to create a block, so we must create a chain to build
// a real block chain

// Essentially, all blockchain has its own struct：a back-linked list, which means the block
// is insert ordered and the next is linked with the lastest block.
// This struct is faster to caculated the hash code.

type Blockchain struct {
	blocks []*Block
}

// Add block into the blockchain
func (bc *Blockchain) AddBlock(data string) {
	preBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

//NewBlockchain :build new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
