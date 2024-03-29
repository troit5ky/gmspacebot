package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type (
	OptionStruct struct {
		Token         string `json:"token"`
		Prefix        string `json:"prefix"`
		IP            string `json:"ip"`
		Status        string `json:"status"`
		StatusCh      bool   `json:"statuschannel"`
		StatusChID    string `json:"statuschannelid"`
		StatusChMsgID string `json:"statuschannelmsgid"`
	}
)

var (
	Option_Example = OptionStruct{
		Token:         "null",
		Prefix:        ".",
		IP:            "127.0.0.1:27015",
		Status:        "Type .help :)",
		StatusCh:      true,
		StatusChID:    "null",
		StatusChMsgID: "null",
	}

	Option *OptionStruct
)

func Init() {
	a, err := ioutil.ReadFile("config/bot.json")

	if err != nil {
		CreateExample()
		os.Exit(0)
	}

	err = json.Unmarshal(a, &Option)

	if err != nil {
		log.Panic(err)
	}
}

func CreateExample() {
	err := os.Mkdir("config", 0755)
	if err != nil {
		log.Fatal(err)
	}

	a, err := json.MarshalIndent(Option_Example, " ", " ")
	if err != nil {
		log.Panic(err)
	}

	ioutil.WriteFile("config/bot.json", a, 0644)
	ioutil.WriteFile("config/bot_example.json", a, 0644)
}

func Rewrite(options OptionStruct) {
	json, err := json.MarshalIndent(options, " ", " ")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("config/bot.json", json, 0644)
}
