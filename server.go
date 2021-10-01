package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

/* type room struct {
	name    string "Go_Chat Box"
	members map[net.Addr]*client
} */

type server struct {
	rooms    string
	commands chan command
}

func newServer() *server {
	return &server{
		//rooms:    string,
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NAME:
			s.name(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		name:     "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) name(c *client, args []string) {
	if len(args) < 2 {
		c.msg("name is required. usage: /name NAME")
		return
	}

	c.name = args[1]
	c.msg(fmt.Sprintf("all right, I will call you %s", c.name))
}

func (s *server) join(c *client, args []string) {

	roomName := args[1]
	r := &room{
		name:    "Go_Chat Box",
		members: make(map[net.Addr]*client),
	}

	r.broadcast(c, fmt.Sprintf("%s joined the room", c.name))

	c.msg(fmt.Sprintf("welcome to %s", roomName))
}

func (s *server) msg(c *client, args []string) {
	if len(args) < 2 {
		c.msg("message is required, usage: /msg MSG")
		return
	}

	msg := strings.Join(args[1:], " ")
	c.broadcast(c, c.name+": "+msg)

}

func (s *server) quit(c *client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())

	//s.quitCurrentRoom(c)

	c.msg("sad to see you go :(")
	c.conn.Close()
}
