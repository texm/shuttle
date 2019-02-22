package bridge

import (
	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"github.com/texm/shuttle/auth"
	"net/url"
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	if (err != nil) {
		return err
	}

	b.IsLoggedIn = true

	return nil
}

func (b *Bridge) LoginWithGoogle() error {
	credentials, err := auth.RetrieveCredentialsThroughOAuth(b.Client)
	if (err != nil) {
		return err
	}

	return b.Login(credentials)
}

func (b *Bridge) SetCredentials(userID string, authToken string) error {
	return nil
}
