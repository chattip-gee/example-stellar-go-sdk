package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/stellar/go/keypair"
)

type KeyPair struct {
	Name    string
	Address string
	Seed    string
}

func main() {

	var pairs []KeyPair
	accountNames := [3]string{"AMT", "GEE", "IYC"}

	for _, element := range accountNames {
		pair, err := keypair.Random()
		if err != nil {
			log.Fatal(err)
		}

		structKeyPair := KeyPair{Name: element, Address: pair.Address(), Seed: pair.Seed()}
		pairs = append(pairs, structKeyPair)
	}

	file, _ := json.MarshalIndent(pairs, "", " ")
	_ = ioutil.WriteFile("pairs.json", file, 0644)
}
