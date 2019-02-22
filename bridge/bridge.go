package bridge

import (
	"fmt"
	"log"
	"net/url"

	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"github.com/texm/shuttle/auth"
)

type PaneState int

const (
	MESSAGE_PANE PaneState = 0
	INPUT_PANE   PaneState = 1
	CHANNEL_PANE PaneState = 2
)

type InterfaceState struct {
	CurViewPanel PaneState
	CurChannel   models.Channel
	CurInput     string
}

type Bridge struct {
	Client     *rest.Client
	IsLoggedIn bool
	uiState    InterfaceState
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools.flnltd.com")
	if err != nil {
		log.Fatalf("bad server url: %s", err)
	}
	brg.Client = rest.NewClient(url, false)

	credentials, err := auth.ReadSavedCredentials()
	if err != nil {
		fmt.Println("failed to read credentials")
		return brg
	}

	err = brg.Login(credentials)
	if err != nil {
		fmt.Println("failed to login")
	}

	return brg
}
