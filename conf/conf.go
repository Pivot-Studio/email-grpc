package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type EmailSenderSettings struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Servername string `json:"servername"`
	Port       int    `json:"port"`
}
type Config struct {
	EmailSenderSettings EmailSenderSettings `json:"EmailSenderSettings"`
}

func ReadSettingsFromFile(settingFilePath string) (config Config) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		fmt.Println(err.Error(), "=======================")
		log.Panic(err)
	}
	return
}
