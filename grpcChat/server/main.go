package main

import "github.com/JayneJacobs/GRPC/grpcChat/chat"

import "io"

import "sync"

import "net"

import "google.golang.org/grpc"

import "fmt"

// Connection creates the Connection STruct
type Connection struct {
	conn chat.Chat_ChatServer
	send chan *chat.ChatMessage
	quit chan struct{}
}

// NewConnection Creates teh connection struct and returns the connection
func NewConnection(conn chat.Chat_ChatServer) *Connection {
	c := &Connection{
		conn: conn,
		send: make(chan *chat.ChatMessage),
		quit: make(chan struct{}),
	}
	go c.start()
	return c
}

// Close feeds teh quit and send channels to close connection
func (c *Connection) Close() error {
	close(c.quit)
	close(c.send)
	return nil
}

// Send sets up teh send channel and passes messages
func (c *Connection) Send(msg *chat.ChatMessage) {
	defer func() {
		recover()
	}()
	c.send <- msg
}

func (c *Connection) start() {
	running := true
	for running {
		select {
		case msg := <-c.send:
			c.conn.Send(msg) // Ignoring the error, they just don't get this message.
		case <-c.quit:
			running = false
		}
	}
}

// GetMessages transmits messages on the broadcast of channels
func (c *Connection) GetMessages(broadcast chan<- *chat.ChatMessage) error {
	for {
		msg, err := c.conn.Recv()
		if err == io.EOF {
			c.Close()
			return nil
		} else if err != nil {
			c.Close()
			return err
		}
		go func(msg *chat.ChatMessage) {
			select {
			case broadcast <- msg:
			case <-c.quit:
			}
		}(msg)
	}
}

// ChatServer fullfills the ChatServer interface
type ChatServer struct {
	broadcast   chan *chat.ChatMessage
	quit        chan struct{}
	connections []*Connection
	connLock    sync.Mutex
}

// NewChatServer creates the broadcast server
func NewChatServer() *ChatServer {
	srv := &ChatServer{
		broadcast: make(chan *chat.ChatMessage),
		quit:      make(chan struct{}),
	}
	go srv.start()
	return srv
}

//Close takes the list of channels and sends messages
func (c *ChatServer) Close() error {
	close(c.quit)
	return nil
}
func (c *ChatServer) start() {
	running := true
	for running {
		select {
		case msg := <-c.broadcast:
			c.connLock.Lock()
			for _, v := range c.connections {
				go v.Send(msg)
			}
			c.connLock.Unlock()
		case <-c.quit:
			running = false
		}
	}
}

// Chat Chat server sets up teh connectitons when a client dials the server
func (c *ChatServer) Chat(stream chat.Chat_ChatServer) error {
	conn := NewConnection(stream)
	c.connLock.Lock()
	c.connections = append(c.connections, conn)
	c.connLock.Unlock()

	err := conn.GetMessages(c.broadcast)

	c.connLock.Lock()
	for i, v := range c.connections {
		if v == conn {
			c.connections = append(c.connections[:i], c.connections[i+1:]...)
		}
	}

	c.connLock.Unlock()
	return err

}

func main() {
	lst, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	srv := NewChatServer()
	chat.RegisterChatServer(s, srv)
	fmt.Println("Serving on port 8081")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}
}
