package main

import (
	"black_mirror/src/api"
	"black_mirror/src/config"

	"github.com/joho/godotenv"
)

func main() {

	db, err := config.ConfigureDb("localhost", 5432, "postgres", "ian", "testdb")

	if db != nil {

		err := godotenv.Load("../.env")

		if err != nil {
			panic(err)
		}

		server := api.NewServer(8000)
		server.Run()

	} else {
		panic(err)
	}

}
