package main

import (
	"fmt"
	"log"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired")
		} else {
			log.Fatal("Config file was found but another error was produced")
		}
	}

	authKey, err := token.AuthKeyFromFile(viper.GetString("P8_FILE_PATH"))
	if err != nil {
		log.Fatal("AuthKey Error:", err)
	}

	token := &token.Token{
		AuthKey: authKey,
		KeyID:   viper.GetString("KEY_ID"),
		TeamID:  viper.GetString("TEAM_ID"),
	}

	notification := &apns2.Notification{}
	notification.DeviceToken = viper.GetString("DEVICE_TOKEN")
	notification.Topic = viper.GetString("BUNDLE_ID")
	notification.Payload = []byte(`{"aps":{"alert":{"body":"Lorem ipsum dolor sit amet consectetur adipiscing elit"},"badge":1,"sound":"default"}}`)

	client := apns2.NewTokenClient(token)
	res, err := client.Push(notification)
	if err != nil {
		log.Fatal("Error:", err)
	}

	if res.Sent() {
		fmt.Println("Sent:", res.ApnsID)
	} else {
		fmt.Printf("Fail: %v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	}
}
