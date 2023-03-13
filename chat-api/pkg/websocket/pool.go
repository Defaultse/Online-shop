package websocket

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Pool struct {
	Register       chan *Client
	Unregister     chan *Client
	Clients        map[string]*websocket.Conn
	ProcessMessage chan Message
	// Services       app.Services
}

func NewPool() *Pool {
	return &Pool{
		Register:       make(chan *Client),
		Unregister:     make(chan *Client),
		Clients:        make(map[string]*websocket.Conn),
		ProcessMessage: make(chan Message),
		// Services:       *services,
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.registerClient(client)
			// fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// for client, _ := range pool.Clients {
			// 	fmt.Println(client)
			// 	client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			// }
			break
		case client := <-pool.Unregister:
			pool.unregisterClient(client)
			// for client, _ := range pool.Clients {
			// 	client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			// }
			break
			// case message := <-pool.ProcessMessage:
			// pool.processMessage(&message)

			// for client, _ := range pool.Clients {
			// 	if err := client.Conn.WriteJSON(message); err != nil {
			// 		fmt.Println(err)
			// 		return
			// 	}
			// }
		}
	}
}

func (pool *Pool) registerClient(client *Client) {
	fmt.Println("Registered client:", client.ID)
	pool.Clients[client.ID] = client.Conn
	fmt.Println("Size of Connection Pool: ", len(pool.Clients))
}

func (pool *Pool) unregisterClient(client *Client) {
	fmt.Println("Unregistered client:", client.ID)
	delete(pool.Clients, client.ID)
	fmt.Println("Size of Connection Pool: ", len(pool.Clients))
}

// func (pool *Pool) processMessage(message *Message) {
// 	result := make(map[string]interface{})
// 	json.Unmarshal([]byte(message.Body), &result)

// 	switch result["type"] {
// 	case 0:
// 		// pool.Services.GetAllLastMessages(message.UserID)
// 	case 1:
// 		panic("Implement!")
// 	}
// }
