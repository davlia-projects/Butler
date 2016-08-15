package main

import (
	"log"

	"github.com/question-service/api"
	"github.com/question-service/config"
)

func main() {
	conf := config.NewConfig()
	s := api.NewServer(conf)
	log.Fatal(s.ListenAndServe())
}
