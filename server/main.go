package main

import (
	"flag"
	"gungnir/config"
	"gungnir/routers"
	"log"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func main() {
	flag.Parse()

	logrus.SetLevel(logrus.DebugLevel)

	r := gin.Default()
	routers.Load(r)
	err := r.Run(":" + config.GetConfig().GetString(config.SYSTEM_LISTEN_PORT))
	if err != nil {
		log.Fatalln("listen port "+config.GetConfig().GetString(config.SYSTEM_LISTEN_PORT)+" failed:", err)
		return
	}
}
