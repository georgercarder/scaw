package main

import (
	. "github.com/georgercarder/scaw/terminal"
)

func main() {
	go NewTerminalSession()
	select{}
}

