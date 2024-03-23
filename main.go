package main

import (
	"fmt"
	"strconv"

	"github.com/peymanzen/blocktest/blockchain"
)

func main(){
	chain := blockchain.InitBlockChain()
	chain.AddBlock("FBAG")
	chain.AddBlock("SBAG")
	chain.AddBlock("TBAG")

	for _, block := range chain.Blocks{
		fmt.Printf("Prev hash: %x\n",block.PrevHash)
		fmt.Printf("Data In Block: %s\n",block.Data)
		fmt.Printf("Hash: %x\n",block.Hash)

		pow := blockchain.newPoW(block)
		fmt.Printf("PoW: %s\n",strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}