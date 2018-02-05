package main

import (
	"encoding/json"
	"flag"
	"os"

	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/server"
)

var configFile = flag.String("config", "config.json", "Config file name")

type config struct {
	Server server.Config `json:"server"`
	DB     data.Config   `jsong:"postgres"`
}

func main() {
	flag.Parse()

	f, err := os.Open(*configFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	conf := config{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}

	db := data.New(conf.DB)
	s := server.New(db, conf.Server)
	s.Run()
}
