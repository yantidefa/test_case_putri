package main

import (
	"log"
	"test_case_putri/config"
	"test_case_putri/routers"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Init()
	routers.InitialRouter()
}
