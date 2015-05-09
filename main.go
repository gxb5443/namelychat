package main

import (
	"log"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.NoRoute(static.Serve("/", static.LocalFile("./public/", false)))
	log.Println("Running Webserver...")
	r.Run(":6000")
}
