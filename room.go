package main

import (
	"net"
)

type room struct {
	name    string "Go_Chat Box"
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {

	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}
