package bridge

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"net/url"
)

type Bridge struct {
	Client *rest.Client
}

func Init(userID string, authToken string) *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools-stg.flnltd.com")
	if (err != nil) {
		// url failed?
	}
	
	client := rest.NewClient(url, false)

	auth := &models.UserCredentials{ID: userID, Token: authToken}
	err = client.Login(auth)
	if (err != nil) {
		// this should never fail?
	}

	brg.Client = client

	return brg
}

