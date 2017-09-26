package main

import (
	_ "github.com/lib/pq"
	"web_site/data"
	"web_site/web_server"
)

func main() {
	db := data.IniDB()
	web_server.RunServer(db)
}
