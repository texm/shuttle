package bridge

import (
	"log"
	"net/url"

	"github.com/texm/shuttle/auth"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
)

type Bridge struct {
	ServerURL url.URL
	Client *rest.Client
	IsLoggedIn bool
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools.flnltd.com")
	if err != nil {
		log.Fatalf("bad server url: %s", err)
	}
	brg.ServerURL = *url

	credentials, err := auth.ReadSavedCredentials()
	if (err == nil) {
		err = brg.Login(credentials)
	}

	return brg
}
