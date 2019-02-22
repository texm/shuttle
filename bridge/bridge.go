package bridge

import (
	"log"
	"net/url"
	"fmt"

	"github.com/texm/shuttle/auth"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
)

type Bridge struct {
	Client *rest.Client
	IsLoggedIn bool
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools.flnltd.com")
	if err != nil {
		log.Fatalf("bad server url: %s", err)
	}
	brg.Client = rest.NewClient(url, false)

	credentials, err := auth.ReadSavedCredentials()
	if (err != nil) {
		fmt.Println("failed to read credentials")
		return brg
	}

	err = brg.Login(credentials)
	if (err != nil) {
		fmt.Println("failed to login")
	}

	return brg
}
