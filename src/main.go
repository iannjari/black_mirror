package main

import (
	"black_mirror/src/api"
	"black_mirror/src/config"
)

func main() {

	db, err := config.GetDB("localhost", 5432, "postgres", "", "testdb")

	if db != nil {

		server := api.NewServer(8000)
		server.Run()

	} else {
		panic(err)
	}

}
