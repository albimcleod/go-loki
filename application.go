package main

import (
	"fmt"

	"github.com/albimcleod/go-loki/core"
	"github.com/albimcleod/go-loki/core/types"
)

func main() {
	var bc types.BlockChain

	fmt.Println("starting go-loki")
	fmt.Println("-----------------")

	lb := core.MakeGenesis()
	bc.Push(*lb)

	fmt.Printf("Genesis Block has been created\n")
	fmt.Printf("Hash: %x\n", lb.Hash())
	fmt.Println("")

	for i := 0; i < 20; i++ {
		lb = core.NextBlock(lb)
		bc.Push(*lb)

		fmt.Printf("Block #%v has been added to the blockchain!\n", lb.Index)
		fmt.Printf("Hash: %x\n", lb.Hash())
		fmt.Println("")
	}
}
