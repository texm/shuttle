package cmd

import (
	"github.com/texm/shuttle/bridge"
	"fmt"
)

func Main(brg *bridge.Bridge) {
	channels, err := brg.GetChannels()
	if (err != nil) {
		fmt.Println("not logged in")
		return
	}

	fmt.Println("channels:\n")
	for i := 0; i < len(channels.Channels); i++ {
		fmt.Println(channels.Channels[i].Name)
	}
}