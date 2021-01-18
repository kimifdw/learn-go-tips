package paxos

type MsgType uint8

const (
	Prepare MsgType = iota
	Promise
	Propose
	Accept
)

// message: 消息
type message struct {
	tp   MsgType
	from int
	to   int
	// number:提议编号
	number int
	// value:提议值
	value string
}
