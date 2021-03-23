package terminal

import (
	"bufio"
	"fmt"

	"os/exec"
	"os"
	"strconv"
	"strings"

	. "github.com/georgercarder/scaw/common"
)


const thisId = "123" // FIXME dummy for now

type tConversation Conversation
type tMessage Message

func NewTerminalSession() {
	conv := new(tConversation)
	// FIXME DUMMY DEMO
	msg1 := &Message{MsgId: "123", TimeStamp: 1, Author: "george", Content: "hello, how are you? This is an example start of a Conversation."}
	conv.Messages = append(conv.Messages, msg1)
	msg2 := &Message{MsgId: "1234", TimeStamp: 2, Author: "bird", Content: "Tweet tweet, I'm a dumb fucking bird. I'm here to chat lol."}
	conv.Messages = append(conv.Messages, msg2)
	msg3 := &Message{MsgId: "1235", TimeStamp: 3, Author: "george", Content: "You bird, are a weird friggin bird. I don't know why I'm talking to a bird. This is soooo strange."}
	conv.Messages = append(conv.Messages, msg3)
	conv.render()
	// ^^ print prev convo
	// loop here
	newMessageCH := make(chan *Message)
	go captureUserInput(newMessageCH)
	for {
		select {
		case m :=<-newMessageCH:
			conv.Messages = append(conv.Messages, m)
			conv.render()
		}
	}
}

func captureUserInput(newMessageCH (chan *Message)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		newMessageCH <- &Message{MsgId: "TODO", TimeStamp: 0, Author: "george", Content: text}
		// TODO
	}
}

func terminalSize() (w, h int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(out), " ")
	if len(arr) != 2 {
		err = fmt.Errorf("terminalSize: could not get w, h.") 
		return
	}
	trimmedW := strings.TrimSpace(arr[0])
	w, err = strconv.Atoi(trimmedW)
	if err != nil {
		return
	}
	trimmedH := strings.TrimSpace(arr[1])
	h, err = strconv.Atoi(trimmedH)
	return
}

func clear(linesPrinted int) (w, h int) {
	w, h, err := terminalSize()
	if err != nil {
		panic(err) // TODO log
	}
	bound := h-linesPrinted
	toPrint := ""
	for i:=0; i<bound; i++ {
		toPrint = "\n"
	}
	fmt.Println(toPrint)
	return
}
