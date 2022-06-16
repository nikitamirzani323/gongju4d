package main

import (
	"log"

	"github.com/nikitamirzani323/gongju4d/routers"
)

func main() {
	app := routers.Init()
	log.Fatal(app.Listen(":7075"))
}
