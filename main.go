//go:generate goagen bootstrap -d github.com/made2591/go-blockchain-go/design

package main

import (

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	//"golang.org/x/tools/cmd/getgo/server"

	"github.com/made2591/go-blockchain-go/app"

	//log "github.com/sirupsen/logrus"

	"strconv"
)

const (
	P2P_PORT = 6001
	HTTP_PORT = 3001
	QUERY_LATEST = 0
	QUERY_ALL = 1
	RESPONSE_BLOCKCHAIN = 2
)

//var SOCKETS = []string{}
//
//func initP2PServer() {
//
//	server = WebSocket.Server({port: P2P_PORT})
//	server.on('connection', ws => initConnection(ws))
//	log.Info("Listening websocket p2p port on: " + P2P_PORT)
//
//}
//
//func initConnection(ws) {
//
//	SOCKETS.push(ws)
//	initMessageHandler(ws)
//	initErrorHandler(ws)
//	write(ws, queryChainLengthMsg())
//
//}

func main() {

	// Create service
	service := goa.New("go-blockchain-go")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "block" controller
	c := NewBlockController(service)
	app.MountBlockController(service, c)
	// Mount "health" controller
	c2 := NewHealthController(service)
	app.MountHealthController(service, c2)
	// Mount "swagger" controller
	c3 := NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "wss" controller
	c4 := NewWssController(service)
	app.MountWssController(service, c4)

	// Start service
	if err := service.ListenAndServe(":"+strconv.Itoa(HTTP_PORT)); err != nil {
		service.LogError("startup", "err", err)
	}

}
