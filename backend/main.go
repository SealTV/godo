package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"bitbucket.org/SealTV/go-site/backend/data"
	"bitbucket.org/SealTV/go-site/backend/service"
)

var configFile = flag.String("config", "config.json", "Config file name")

type config struct {
	Service service.Config `json:"server"`
	DB      data.Config    `jsong:"postgres"`
}

func main() {
	flag.Parse()
	log.Println("Start")

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
	s := service.New(db, conf.Service)
	s.Run()
}
