package bridge

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Billz95/Rocket.Chat.Go.SDK/realtime"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"github.com/texm/shuttle/auth"
)

type Bridge struct {
	ServerURL      url.URL
	IsLoggedIn     bool
	Client         *rest.Client
	RealtimeClient *realtime.Client
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools.flnltd.com")
	if err != nil {
		log.Fatalf("bad server url: %s", err)
	}
	brg.Client = rest.NewClient(url, false)
	brg.RealtimeClient, err = realtime.NewClient(url, false)
	if err != nil {
		log.Fatalf("couldn't make new realtime client: %s", err)
	}
	credentials, err := auth.ReadSavedCredentials()
	if err != nil {
		fmt.Printf("failed to read credentials: %s\n", err)
		return brg
	}
	err = brg.Login(credentials)
	if err != nil {
		fmt.Printf("failed to login: %s\n", err)
	}

	return brg
}
