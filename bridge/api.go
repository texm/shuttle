package bridge

import (
	"github.com/RocketChat/Rocket.Chat.Go.SDK/rest"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"

	"net/url"
)

func (b *Bridge) SendMessage(msg models.PostMessage) {
	// send
}

func (b *Bridge) GetChannels() (*rest.ChannelsResponse, error) {
	return b.Client.GetPublicChannels()
}

func (b *Bridge) GetJoinedChannels(params url.Values) (*rest.ChannelsResponse, error) {
	return b.Client.GetJoinedChannels(params)
}

func (b *Bridge) Search(params url.Values) (*models.Spotlight, error) {
	return b.Client.GetSpotlight(params)
}