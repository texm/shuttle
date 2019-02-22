package main

import "flag"

import "./cmd"
import "./ui"

func main() {
	noUIptr := flag.Bool("noui", false, "use UI or command")
	flag.Parse()

	if (*noUIptr) {
		cmd.Main()
	} else {
		ui.Main()
	}
}
