package bridge

import (
	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
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

func (b *Bridge) Login(c *rest.Client) error {
	credential, _ := ReadSavedCredential()
	_ = c.LoginViaGoogle(&credential)
	SaveCredential(&credential)
	return nil
}

func ReadSavedCredential() (models.UserCredentials, error) {
	pwd, _ := os.Getwd()
	fmt.Println(pwd + "/.credential")
	data, err := ioutil.ReadFile(pwd + "/.credential")
	fmt.Println(data)
	fmt.Println(err)
	if err != nil {
		return models.UserCredentials{}, err
	}

	dat := models.UserCredentials{}
	json.Unmarshal(data, &dat)
	return models.UserCredentials{ID: dat.ID, Token: dat.Token}, nil
}

func SaveCredential(credentials *models.UserCredentials) error {
	pwd, _ := os.Getwd()
	se, _ := json.Marshal(credentials)
	err := ioutil.WriteFile(pwd+"/.credential", se, 0644)
	return err
}

func AppLogin(c *rest.Client) error {
	credential, _ := ReadSavedCredential()
	fmt.Println("===========================")
	fmt.Println(credential)
	fmt.Println("===========================")
	_ = c.LoginViaGoogle(&credential)
	SaveCredential(&credential)
	return nil
}

func (b *Bridge) SetCredentials(userID string, authToken string) error {
	return nil
}