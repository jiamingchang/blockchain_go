package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

var cli = CLI{}

func getBalance(context *gin.Context){
	getBalanceAddress := context.PostForm("getBalanceAddress")
	nodeID := context.PostForm("nodeID")
	balance, ok := cli.getBalance(getBalanceAddress, nodeID)

	context.JSON(200, gin.H{
		"isSuccess": ok,
		"balance": balance,
	})
}
func createBlockchain(context *gin.Context){
	createBlockchainAddress := context.PostForm("getBalanceAddress")
	nodeID := context.PostForm("nodeID")
	ok := cli.createBlockchain(createBlockchainAddress, nodeID)
	context.JSON(200, gin.H{
		"isSuccess": ok,
	})
}
func createWallet(context *gin.Context){
	nodeID := context.PostForm("nodeID")
	address := cli.createWallet(nodeID)
	context.JSON(200, gin.H{
		"address": address,
	})
}
func listAddresses(context *gin.Context){
	nodeID := context.PostForm("nodeID")
	addresses := cli.listAddresses(nodeID)
	context.JSON(200, gin.H{
		"addresses": addresses,
	})
}
func printChain(context *gin.Context){
	nodeID := context.PostForm("nodeID")
	blocks := cli.printChain(nodeID)
	context.JSON(200, gin.H{
		"blocks": blocks,
	})
}
func reindexUTXO(context *gin.Context){
	nodeID := context.PostForm("nodeID")
	count, ok := cli.reindexUTXO(nodeID)
	context.JSON(200, gin.H{
		"isSuccess": ok,
		"count": count,
	})
}
func send(context *gin.Context){
	sendFrom := context.PostForm("sendFrom")
	sendTo := context.PostForm("sendTo")
	sendAmount := context.PostForm("sendAmount")
	amount,_ := strconv.Atoi(sendAmount)
	sendMine := context.PostForm("sendMine")
	mine,_ := strconv.ParseBool(sendMine)
	nodeID := context.PostForm("nodeID")
	ok := cli.send(sendFrom, sendTo, amount, nodeID, mine)
	context.JSON(200, gin.H{
		"isSuccess": ok,
	})
}
func startNode(context *gin.Context){
	startNodeMiner := context.PostForm("startNodeMiner")
	nodeID := context.PostForm("nodeID")
	ok := cli.startNode(nodeID, startNodeMiner)
	context.JSON(200, gin.H{
		"isSuccess": ok,
	})
}
func storedTransaction(context *gin.Context){
	storedTransactionFrom := context.PostForm("storedTransactionFrom")
	storedTransactionMine := context.PostForm("storedTransactionMine")
	mine,_ := strconv.ParseBool(storedTransactionMine)
	nodeID := context.PostForm("nodeID")
	ok := cli.storedTransaction(storedTransactionFrom, nodeID, mine)
	context.JSON(200, gin.H{
		"isSuccess": ok,
	})
}
func getdataamount(context *gin.Context){
	getdataaddress := context.PostForm("getdataaddress")
	nodeID := context.PostForm("nodeID")
	count ,ok := cli.getdataamount(getdataaddress, nodeID)
	context.JSON(200, gin.H{
		"isSuccess": ok,
		"count": count,
	})
}
