package main

import (
	"bufio"
	"fmt"

	"os/exec"
	"os"
	"strconv"
	"strings"
)

const thisId = "123"

type message struct {
	msgId string // TODO
	timeStamp uint32
	author string // TODO struct
	authorId string
	content string
}

type conversation struct {
	messages []*message
}

func main() {
	conv := new(conversation)
	msg1 := &message{msgId: "123", timeStamp: 1, author: "george", content: "hello, how are you? This is an example start of a conversation."}
	conv.messages = append(conv.messages, msg1)
	msg2 := &message{msgId: "1234", timeStamp: 2, author: "bird", content: "Tweet tweet, I'm a dumb fucking bird. I'm here to chat lol."}
	conv.messages = append(conv.messages, msg2)
	msg3 := &message{msgId: "1235", timeStamp: 3, author: "george", content: "You bird, are a weird friggin bird. I don't know why I'm talking to a bird. This is soooo strange."}
	conv.messages = append(conv.messages, msg3)
	conv.render()
	// ^^ print prev convo
	// loop here
	newMessageCH := make(chan *message)
	go captureUserInput(newMessageCH)
	for {
		select {
		case m :=<-newMessageCH:
			conv.messages = append(conv.messages, m)
			conv.render()
		}
	}
	select{}
}

func captureUserInput(newMessageCH (chan *message)) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		newMessageCH <- &message{msgId: "TODO", timeStamp: 0, author: "george", content: text}
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

func clear() (w, h int) {
	w, h, err := terminalSize()
	if err != nil {
		panic(err) // TODO log
	}
	for i:=0; i<h; i++ {
		fmt.Println("\n")
	}
	return
}

func (c *conversation) render() {
	w, h := clear()
	for _, m := range c.messages {
		m.print(w, h)
	}
	fmt.Print(" > ") // cursor
}

func (m *message) print(w, h int) {
	fmt.Printf("-%s-\n  %s\n\n", m.author, m.content)
}
