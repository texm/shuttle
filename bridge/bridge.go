package bridge

import (
	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"net/url"
)

type Bridge struct {
	Client *rest.Client
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools-stg.flnltd.com")
	if err != nil {
		// url failed?
	}

	client := rest.NewClient(url, false)
	err = brg.AppLogin(client)

	// google oauth
	auth := &models.UserCredentials{}

	auth.ID = "DETgMpoxPtE22Ap7Y"
	auth.Token = "9QLabFKs_-gVTqOeaso6PbsskWmsuxY2yXA6AL-d_Kx"

	err = client.Login(auth)
	if err != nil {
		// this should never fail?
	}

	brg.Client = client
	return brg
}
