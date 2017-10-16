package main

func main() {

	bc := NewBlockchain()
	bc.AddBlock("send 1 btc to me")
	bc.AddBlock("send 2 more btc to me")

	for _, block := range bc.blocks {
		println("pre. hash:\t%x\n", block.PrevBlockHash)
		println("data:\t%s\n", block.Data)
		println("hash:\t%x\n", block.Hash)
		println()
	}

}
