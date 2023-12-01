package main

import (
	"log"

	api "micah.wiki/pandora/kitex_gen/api/pandora"
)

func main() {
	svr := api.NewServer(new(PandoraImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
