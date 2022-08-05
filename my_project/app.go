package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/config"
	"my_project/handlers"
	"my_project/logger"
	"net"
	"os"
	"path"
	"path/filepath"
)
import _ "github.com/lib/pq"

func main() {
	//database.Init()
	//database.GetDB()
	////database.Selected()
	//database.Delete()
	////database.Created()

	logs := logger.GetLogger()
	logs.Info("запуск сервера")

	router := gin.Default()
	conf := config.GetConfig()
	start(conf)
	router.GET("/home", handlers.Home)
	router.Run(":8081")

}
func start(conf *config.Config) {
	// при смене тайпа в конфиге с порта на сокет все равно идет в ветку tcp
	log := logger.GetLogger()
	log.Info("запуск приложения")

	var listener net.Listener

	if conf.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}

		log.Info("создаем сокет")
		socketPath := path.Join(appDir, "app.sock")
		//log.Debugf("socket path %s", socketPath)
		log.Info("слушаем unix сокет")
		listener, err = net.Listen("unix", socketPath)
		log.Infof("слушаем sock %s", socketPath)
		if err != nil {
			panic(err)
		}
	} else {
		log.Info("слушаем tcp")
		var err error
		listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Listen.BindIp, conf.Listen.Port))
		log.Infof("слушаем порт %s:%s", conf.Listen.BindIp, conf.Listen.Port)
		if err != nil {
			panic(err)
			log.Fatal(listener)
		}
	}
	//log.Infof("слушаем порт %s:%s", conf.Listen.BindIp, conf.Listen.Port)

}
