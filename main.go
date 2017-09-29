package main

import (
	"./data"
	"./server"
)

func main() {
	db := data.IniDB()
	server.RunServer(db)
}
