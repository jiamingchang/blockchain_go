package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

func (cli *CLI) storedTransaction(from string, nodeID string, mineNow bool) bool {
	if !ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
		return false
	}

	bc := NewBlockchain(nodeID)
	UTXOSet := UTXOSet{bc}
	defer bc.db.Close()

	randData := make([]byte, 20)
	_, err := rand.Read(randData)
	if err != nil {
		log.Panic(err)
		return false
	}

	data := fmt.Sprintf("%x", randData)

	txin := TXInput{[]byte{}, -1, nil, []byte(data)}
	txout := NewTXOutput(0, from)
	tx := Transaction{nil, "store", []TXInput{txin}, []TXOutput{*txout}}
	pubKeyHash := Base58Decode([]byte(from))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	tx.ID = pubKeyHash

	if mineNow {
		cbTx := NewCoinbaseTX(from, "")
		txs := []*Transaction{cbTx, &tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		sendTx(knownNodes[0], &tx)
	}

	fmt.Println("Success!")
	return true
}
