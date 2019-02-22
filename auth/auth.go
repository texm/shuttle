package auth

import (
	"encoding/json"
	"fmt"
	"github.com/Billz95/Rocket.Chat.Go.SDK/models"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"io/ioutil"
	"os"
)

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
