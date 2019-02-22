package main

import (
	"flag"

	"github.com/texm/shuttle/bridge"
	"github.com/texm/shuttle/cmd"
	"github.com/texm/shuttle/ui"
)

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
