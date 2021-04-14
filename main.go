package main

import "github.com/xenobyter/xbVorrat/api"


func main() {
	router := api.SetupRouter()
	router.Run()
}

