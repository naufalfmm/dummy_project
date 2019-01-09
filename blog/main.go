package main

import (
	"log"
)

func main() {
	router := NewRouter()

	log.Fatal(router.Run(":8080"))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
