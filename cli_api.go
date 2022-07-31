package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	// 取消ws跨域校验
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var mux sync.Mutex
//私聊客户端链接池
var lv1client = make(map[string]*websocket.Conn)
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
	cli.startNode(nodeID, startNodeMiner)
	context.JSON(200, gin.H{
		"isSuccess": true,
	})
}
func stopNode(context *gin.Context){
	nodeID := context.PostForm("nodeID")
	ok := cli.stopNode(nodeID)
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
func Wshandlesendmessage(context *gin.Context)  {
	//获取链接客户端id
	postid:=context.Query("id")

	log.Println(postid+" websocket客户端连接")
	//websocket传输
	wshandlewrjfunc(context.Writer, context.Request, postid)
}



func wshandlewrjfunc(w http.ResponseWriter, r *http.Request, postid string)  {
	var conn *websocket.Conn	//websocket客户端
	var message lv1firstchat             //将数据解析的格式
	var err error
	cnt:=0
	//将http请求升级为webscoket链接
	conn, err = wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 把与客户端的链接添加到客户端链接池中
	lv1addClient(postid, conn)
	// 设置客户端关闭ws链接回调函数
	conn.SetCloseHandler(func(code int, text string) error {
		lv1deleteClient(postid)
		log.Println(code)
		return nil
	})
	for{
		log.Println("开始接受并解析数据(json序列化)")
		err=conn.ReadJSON(&message)
		log.Println("message：",message)
		if err!=nil{
			log.Println("出错")
			cnt++
			log.Println(err)
			lv1deleteClient(postid)
		}
		if cnt==3{
			break
		}
		//转发数据给web
		if message.Postuserid!=""{
			webcon,exist:=lv1getClient(message.Receiveuserid)
			if exist{
				webcon.WriteJSON(&message)
			}else {
				err:=conn.WriteMessage(websocket.TextMessage,[]byte(message.Receiveuserid+"不在线"))
				if err!=nil{
					log.Println(err)
				}
				log.Println(message.Receiveuserid+"不在线")
			}


		}
		message=lv1firstchat{}


	}



}
func lv1addClient(id string, conn *websocket.Conn) {
	mux.Lock()
	lv1client[id] = conn
	mux.Unlock()
}

func lv1getClient(id string) (conn *websocket.Conn, exist bool) {
	mux.Lock()
	conn, exist = lv1client[id]
	mux.Unlock()
	return
}

func lv1deleteClient(id string) {
	mux.Lock()
	delete(lv1client, id)
	log.Println(id + " websocket私聊客户端退出")
	mux.Unlock()
}

type lv1firstchat struct {
	Postuserid string
	Message string
	Messagetype int
	Receiveuserid string
}
