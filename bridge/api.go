package bridge

import (
	"fmt"
	"net/url"

	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"github.com/texm/shuttle/auth"
)

func (b *Bridge) SendMessage(msg string, channel *models.Channel) error {
	// send`

	return b.Client.Send(channel, msg)
}

func (b *Bridge) GetMessages(channel *models.Channel, page *models.Pagination) ([]models.Message, error) {
	// cache

	return b.Client.GetMessages(channel, page)
}

func (b *Bridge) GetChannels() (*rest.ChannelsResponse, error) {
	// cache
	return b.Client.GetPublicChannels()
}

func (b *Bridge) GetJoinedChannels(params url.Values) (*rest.ChannelsResponse, error) {
	// cache
	return b.Client.GetJoinedChannels(params)
}

func (b *Bridge) Search(params url.Values) (*models.Spotlight, error) {
	// cache
	return b.Client.GetSpotlight(params)
}

func (b *Bridge) Login(credentials *models.UserCredentials) error {
	err := b.Client.Login(credentials)
	if err != nil {
		return err
	}
	_, err = b.RealtimeClient.Login(credentials)
	if err != nil {
		return fmt.Errorf("couldn't login to realtime client: %s", err)
	}

	b.IsLoggedIn = true

	return nil
}

func (b *Bridge) StreamMessages(channel *models.Channel) (chan models.Message, error) {
	msgCh := make(chan models.Message, 100) // just an arbitrary buffer size
	err := b.RealtimeClient.SubscribeToMessageStream(channel, msgCh)
	if err != nil {
		close(msgCh)
		return nil, fmt.Errorf("error subscribing to message stream: %s", err)
	}
	return msgCh, nil
}

func (b *Bridge) LoginWithGoogle() error {
	credentials, err := auth.RetrieveCredentialsThroughOAuth(b.Client)
	if err != nil {
		return err
	}

	return b.Login(credentials)
}

func (b *Bridge) SetCredentials(userID string, authToken string) error {
	return nil
}

func (b *Bridge) SetPaneState(s PaneState) error {
	// cache
	b.uiState.CurViewPanel = s
	return nil
}

func (b *Bridge) SetCurChannel(c models.Channel) error {
	b.uiState.CurChannel = c
	return nil
}

func (b *Bridge) setCurInput(s string) error {
	b.uiState.CurInput = s
	return nil
}

func (b *Bridge) getCurChannel() models.Channel {
	return b.uiState.CurChannel
}

func (b *Bridge) getPaneState() PaneState {
	return b.uiState.CurViewPanel
}
