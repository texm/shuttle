package bridge

import (
	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"

	"net/url"
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
