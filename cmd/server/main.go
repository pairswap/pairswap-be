package main

import (
	"github.com/sisu-network/pairswap-be/src/handler"
)

func main() {
	app := handler.NewApp()
	app.Start()
	select {}
}
