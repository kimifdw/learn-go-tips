package main

import (
	"errors"
	"sync"
	"time"
)

const (
	nodeBits  uint8 = 10
	stepBits  uint8 = 12
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift       = nodeBits + stepBits
	nodeShift       = stepBits
)

var Epoch int64 = 1288834974657

type ID int64

type Node struct {
	mu        sync.Mutex
	timestamp int64
	node      int64
	step      int64
}

func NewNode(node int64) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 0 and 1023")
	}

	return &Node{
		timestamp: 0,
		node:      node,
		step:      0,
	}, nil
}

func (n *Node) Generate() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now().UnixNano() / 1e6
	if n.timestamp == now {
		n.step++

		if n.step > stepMax {
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		n.step = 0
	}

	n.timestamp = now
	result := ID((now-Epoch)<<timeShift | (n.node << nodeShift) | (n.step))

	return result
}

//func main() {
//
//   node, err := NewNode(1)
//   if err != nil {
//       fmt.Println(err)
//       return
//   }
//
//   ch := make(chan ID)
//   count := 10000
//
//   for i := 0; i < count; i++ {
//       go func() {
//           id := node.Generate()
//           ch <- id
//       }()
//   }
//
//   defer close(ch)
//
//   m := make(map[ID]int)
//   for i := 0; i < count; i++ {
//       id := <-ch
//       fmt.Printf("id:%d\n", id)
//       _, ok := m[id]
//       if ok {
//           fmt.Printf("ID is not unique!\n")
//           return
//       }
//       m[id] = i
//   }
//
//   fmt.Println("All ", count, " snowflake ID generate successed!\n")
//}
