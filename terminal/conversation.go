package terminal

import (
	"fmt"
)

func (c *tConversation) render() {
	w, h := clear(c.LinesPrinted)
	lineSum := 0
	idx := 0
	// optimize for terminal height
	for i:=len(c.Messages)-1; i>-1; i-- {
		msg := (*tMessage)(c.Messages[i])
		lineSum = msg.length(w)
		if lineSum > h {
			idx = i
			break
		}
	}
	linesPrinted := 0
	// print
	for i:=idx; i< len(c.Messages); i++ {
		msg := (*tMessage)(c.Messages[i])
		linesPrinted = msg.print(w, h)
	}
	fmt.Print(" > ") // cursor
	linesPrinted += 1 // cursor
	c.LinesPrinted = linesPrinted
}

func (m *tMessage) length(w int) (lineSum int) {
	lineSum = len(m.Content) / w
	return
}

func (m *tMessage) print(w, h int) (linesPrinted int) {
	fmt.Printf("-%s-\n  %s\n\n", m.Author, m.Content)
	linesPrinted = 1 // author line
	linesPrinted += len(m.Content) / w
	return
}
