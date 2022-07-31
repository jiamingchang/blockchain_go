package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) bool {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
		return false
	}
	bc := CreateBlockchain(address, nodeID)
	if bc == nil{
		return false
	}
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
	return true
}
