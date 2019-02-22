package main

import "flag"

import "./bridge"
import "./cmd"
import "./ui"

func main() {
	noUIptr := flag.Bool("noui", false, "use UI or command")
	flag.Parse()

	bridge.Init()

	if (*noUIptr) {
		cmd.Main()
	} else {
		ui.Main()
	}
}
