package main

import (
	"fmt"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	serverPort := config.GetEnvString("PORT", "18080")
	r := gin.Default()
	log.Fatal(r.Run(fmt.Sprintf(":%s", serverPort)))

}
