package bridge

import (
	"fmt"
	// "io/ioutil"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/realtime"
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
	Client         *rest.Client
	IsLoggedIn     bool
	uiState        InterfaceState
	ServerURL      url.URL
	RealtimeClient *realtime.Client
	User           *UserInfoStruct
}

type UserInfoStruct struct {
	Name     string `json:"name"`
	UserName string `json:"username"`
	ID       string `json:"_id"`
}

func (o UserInfoStruct) OK() error {
	return nil
}

func Init() *Bridge {
	brg := &Bridge{}

	url, err := url.Parse("https://chat.tools.flnltd.com")
	if err != nil {
		log.Fatalf("bad server url: %s", err)
	}
	brg.Client = rest.NewClient(url, false)

	log.SetOutput(ioutil.Discard)

	brg.RealtimeClient, err = realtime.NewClient(url, false)
	if err != nil {
		log.Fatalf("couldn't make new realtime client: %s", err)
	}

	credentials, err := auth.ReadSavedCredentials()
	if err != nil {
		fmt.Println("failed to read credentials")
		return brg
	}
	err = brg.Login(credentials)
	if err != nil {
		fmt.Printf("failed to login: %s\n", err)
	} else {
		res, _ := brg.GetUserInfo()
		brg.User = res
	}

	return brg
}

func (b *Bridge) GetInterfaceState() *InterfaceState {
	chanResponse, err := b.GetJoinedChannels(url.Values{})
	if err != nil {
		fmt.Println(err)
	}
	var channel models.Channel

	for i := 0; i < len(chanResponse.Channels); i++ {
		if chanResponse.Channels[i].Name == "shuttle-test" {
			channel = chanResponse.Channels[i]
			break
		}
	}

	if channel.Name == "" {
		channel = chanResponse.Channels[0]
	}

	state := &InterfaceState{CurInput: "", CurChannel: channel, CurViewPanel: INPUT_PANE}

	return state
}
