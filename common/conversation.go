package common

type Message struct {
	MsgId string // TODO
	TimeStamp uint32
	Author string // TODO struct
	AuthorId string
	Content string
}

type Conversation struct {
	LinesPrinted int
	Messages []*Message
}

