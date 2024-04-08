package main

import (
	"black_mirror/src/config"
	"net/http"
)

func main() {

	db, err := config.GetDB("localhost", 5432, "postgres", "", "testdb")

	if db != nil {

		http.ListenAndServe("localhost:8000", http.NewServeMux())

	} else {
		panic(err)
	}

}
