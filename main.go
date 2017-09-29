package main

import (
	"go-site/data"
	"go-site/server"
)

func main() {
	db := data.IniDB()
	server.RunServer(db)
}
