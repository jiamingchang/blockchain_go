package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type CLI struct {}

func (cli *CLI) Run(){
	router := gin.Default()
	go start_init()

	v1 := router.Group("/v1")
	{
		v1.POST("/getbalance", getBalance)
		v1.POST("/createblockchain", createBlockchain)
		v1.POST("/createwallet", createWallet)
		v1.POST("/listaddresses", listAddresses)
		v1.POST("/printchain", printChain)
		v1.POST("/reindexutxo", reindexUTXO)
		v1.POST("/send", send)
		v1.POST("/startnode", startNode)
		v1.POST("/stopnode", stopNode)
		v1.POST("/storedTransaction", storedTransaction)
		v1.POST("/getdataAmount", getdataamount)
		v1.GET("/sendmessage",Wshandlesendmessage)
	}
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}

func start_init(){
	address:=cli.createWallet("3000")
	cli.createBlockchain(address,"3000")
	cli.startNode("3000",address)
}