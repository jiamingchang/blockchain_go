package main

import (
	"fmt"
	"strconv"
)

func (cli *CLI) printChain(nodeID string) []*Block {
	bc := NewBlockchain(nodeID)
	defer bc.db.Close()

	bci := bc.Iterator()
	var blocks []*Block
	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		blocks = append(blocks, block)
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return blocks
}
