package main

import (
	"github.com/goadesign/goa"
	"github.com/gorilla/websocket"
	"github.com/made2591/go-blockchain-go/app"
	"encoding/json"
)

// WssController implements the wss resource.
type WssController struct {
	*goa.Controller
}

// NewWssController creates a wss controller.
func NewWssController(service *goa.Service) *WssController {
	return &WssController{Controller: service.NewController("WssController")}
}

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Content: "A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.id, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
//
//func initMessageHandler(ws *websocket.Conn) {
//
//		var message = JSON.parse(ws.ReadMessage());
//		console.log('Received message' + JSON.stringify(message));
//
//		switch (message.type) {
//			case MessageType.QUERY_LATEST:
//				write(ws, responseLatestMsg())
//				break
//			case MessageType.QUERY_ALL:
//				write(ws, responseChainMsg())
//				break
//			case MessageType.RESPONSE_BLOCKCHAIN:
//				handleBlockchainResponse(message)
//				break
//		}
//	})
//}

// Wsn runs the wsn action.
func (c *WssController) Wsn(ctx *app.WsnWssContext) error {
	// WssController_Wsn: start_implement

	go manager.start()

	// WssController_Wsn: end_implement
	return nil
}
