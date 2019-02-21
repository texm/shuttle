package main

import "flag"

func main() {
	noUIptr := flag.Bool("noui", false, "use UI or command")
	flag.Parse()

	if (*noUIptr) {
		cmd_main()
	} else {
		ui_main()
	}
}
