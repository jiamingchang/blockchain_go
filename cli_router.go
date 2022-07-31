package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type CLI struct {}

func (cli *CLI) Run(){
	router := gin.Default()

	address:=cli.createWallet("8080")
	cli.createBlockchain(address,"8080")
	cli.startNode("8080",address)
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
		v1.POST("/storedTransaction", storedTransaction)
		v1.POST("/getdataAmount", getdataamount)
		v1.GET("/sendmessage",Wshandlesendmessage)
	}
	err := router.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
