package paxos

import (
	"log"
	"time"
)

// network:网络实现
type network interface {
	// send 发送消息
	send(m message)
	// recv 接收消息
	recv(timeout time.Duration) (message, bool)
}

type Network struct {
	queue map[int]chan message
}

// newNetwork 初始化
func newNetwork(nodes ...int) *Network {
	pn := &Network{
		queue: make(map[int]chan message, 0),
	}

	for _, a := range nodes {
		pn.queue[a] = make(chan message, 1024)
	}
	return pn
}

// send 发送消息
func (net *Network) send(m message) {
	log.Printf("net: send %+v", m)
	net.queue[m.to] <- m
}

// recvFrom：接收消息
func (net *Network) recvFrom(from int, timeout time.Duration) (message, bool) {
	select {
	case m := <-net.queue[from]:
		log.Printf("net: recv %+v", m)
		return m, true
	case <-time.After(timeout):
		return message{}, false
	}
}

type nodeNetwork struct {
	id int
	*Network
}

func (n *nodeNetwork) send(m message) {
	n.Network.send(m)
}

func (n *nodeNetwork) recv(timeout time.Duration) (message, bool) {
	return n.recvFrom(n.id, timeout)
}

func (net *Network) nodeNetwork(id int) *nodeNetwork {
	return &nodeNetwork{
		id:      id,
		Network: net,
	}
}
