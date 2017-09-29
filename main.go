package main

import (
	"bitbucket.org/SealTV/go-site/data"
	"bitbucket.org/SealTV/go-site/server"
)

func main() {
	db := data.IniDB()
	server.RunServer(db)
}
