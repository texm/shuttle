package auth

import (
	"os"
	"encoding/json"
	"io/ioutil"
	"errors"

	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
)

func ReadSavedCredentials() (*models.UserCredentials, error) {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "/.credential")

	credentials := &models.UserCredentials{}
	if err != nil {
		return credentials, err
	}

	json.Unmarshal(data, &credentials)

	if (credentials.ID == "" || credentials.Token == "") {
		return credentials, errors.New("failed to read tokens")
	}

	return credentials, nil
}

func SaveCredentials(credentials *models.UserCredentials) error {
	pwd, _ := os.Getwd()

	data := map[string]string{"ID": credentials.ID, "Token": credentials.Token}

	se, _ := json.Marshal(data)
	err := ioutil.WriteFile(pwd+"/.credential", se, 0644)
	return err
}

func RetrieveCredentialsThroughOAuth(client *rest.Client) (*models.UserCredentials, error) {
	credentials := &models.UserCredentials{}

	err := client.LoginViaGoogle(credentials)
	if (err != nil) {
		return credentials, err
	}

	SaveCredentials(credentials)

	return credentials, nil
}