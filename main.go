package main

import (
	"web_site/data"
	"web_site/server"
)

func main() {
	db := data.IniDB()
	server.RunServer(db)
}
