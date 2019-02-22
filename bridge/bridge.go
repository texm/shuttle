package bridge

import (
	"log"
	"net/url"

	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
)

type Bridge struct {
	Client *rest.Client
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools.flnltd.com")
	if err != nil {
		log.Fatalf("bad server url: %s", err)
	}

	client := rest.NewClient(url, false)
	err = brg.Login(client)

	if err != nil {
		log.Fatalf("couldn't login: %s", err)
	}

	brg.Client = client
	return brg
}
