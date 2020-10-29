package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	ADRESSE_BROKER_MQTT string `json:"ADRESSE_BROKER_MQTT"`
	PORT_BROKER_MQTT string `json:"PORT_BROKER_MQTT"`
	NIVEAU_DE_QOS string `json:"NIVEAU_DE_QOS"`
	ID_CLIENT string `json:"ID_CLIENT"`
	MOT_DE_PASSE string `json:"MOT_DE_PASSE"`
}

const nomFichier string = "./external/sub/ressources/config.json"

func GetConfig() Configuration {
	jsonFile, err := os.Open(nomFichier)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config Configuration

	json.Unmarshal(byteValue, &config)

	return config
}
