package main

import (
	"bytes"
	"fmt"
	"log"
)

func (cli *CLI) getdataamount(address, nodeID string)(int, bool) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
		return 0, false
	}
	var count int
	pubKeyHash := Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]

	bc := NewBlockchain(nodeID) //返回区块链中最后一个区块
	defer bc.db.Close()
	bci := bc.Iterator()
	for {
		block := bci.Next()
		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, pubKeyHash) == 0 && tx.Form == "store" {
				count += 100
			}
		}
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	fmt.Printf("The address: \033[1;32;40m%s\033[0m has stored \033[1;31;40m%d\033[0m MB data \n", address, count)
	return count, true
}
