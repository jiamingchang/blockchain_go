package main

import (
	"context"
	"fmt"
	"log"
)

type cont struct {
	nodeID string
	ctx    context.Context
	cancel context.CancelFunc
}
var conts []cont

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	c := cont{
		nodeID: nodeID,
		ctx: ctx,
		cancel: cancel,
	}
	conts = append(conts, c)
	go StartServer(ctx, nodeID, minerAddress)
}

func (cli *CLI) stopNode(nodeID string) bool {
	for i, c := range conts{
		if c.nodeID == nodeID {
			c.cancel()
			conts = append(conts[:i], conts[i+1:]...)
			log.Println("停止节点:"+c.nodeID)
			return true
		}
	}
	return false
}