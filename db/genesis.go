package db

import (
	"encoding/json"
	"io/ioutil"
)

type genesis struct {
	Balances map[Account]uint `json:"balances"`
}

type genesistime struct {
	GenesisTime string `json:"genesis_time"`
}

func loadGenesis(path string) (genesis, error) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		return genesis{}, err
	}

	var loadedGenesis genesis
	err = json.Unmarshal(content, &loadedGenesis)
	var time genesistime

	err = json.Unmarshal(content, &time)

	if err != nil {
		return genesis{}, err
	}

	return loadedGenesis, nil
}
