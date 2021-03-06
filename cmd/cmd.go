package cmd

import (
	"fmt"
	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/texm/shuttle/bridge"
)

func Main(brg *bridge.Bridge) {
	channels, err := brg.GetChannels()
	if err != nil {
		fmt.Println("not logged in")
		return
	}

	var testChan *models.Channel
	for i := 0; i < len(channels.Channels); i++ {
		if channels.Channels[i].Name == "shuttle-test" {
			testChan = &channels.Channels[i]
			break
		}
	}

	if testChan == nil {
		fmt.Println("not found")
		return
	}

	err = brg.SendMessage("test", testChan)
	if err == nil {
		fmt.Println(err)
	}
}
