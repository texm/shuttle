package main

import (
	// "encoding/json"
	// "fmt"
	// "github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"fmt"
	"github.com/Billz95/Rocket.Chat.Go.SDK/rest"
	"net/url"
)

func main() {
	//
	client := rest.NewClient(&url.URL{Host: "chat.tools.flnltd.com:80"}, false)
	AppLogin(client)
	// credential := models.UserCredentials{}

	// se, _ := json.Marshal(models.UserCredentials{ID: "abc", Token: "kkk", Email: "mail"})
	// fmt.Println(se)
	// fmt.Println(string(se))
	// dat := models.UserCredentials{}
	// json.Unmarshal(se, &dat)
	// fmt.Println(dat)

	// client.LoginViaGoogle(&credential)
	// fmt.Println(credential)
	res, err := client.GetPublicChannels()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
