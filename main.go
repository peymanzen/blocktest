package main

import (
	"github.com/peymanzen/blocktest/blockchain"
	"fmt"
)

func main(){
	chain := InitBlockChain()
	chain.AddBlock("FBAG")
	chain.AddBlock("SBAG")
	chain.AddBlock("TBAG")

	for _, block := range chain.blocks{
		fmt.Printf("Prev hash: %x\n",block.PrevHash)
		fmt.Printf("Data In Block: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)
	}
}