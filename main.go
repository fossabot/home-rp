package main

import (
	"github.com/just1689/home-rp/model"
	"github.com/just1689/home-rp/server"
	"github.com/just1689/home-rp/util"
	"github.com/sirupsen/logrus"
)

const (
	EnvKeyListenAddr     = "ENV_KEY_LISTEN_ADDRESS"
	DefaultListenAddress = ":8000"
	EnvKeyRoutes         = "ENV_KEY_ROUTES"
	DefaultRoutes        = "" //HOST=FULL_URL,HOST=FULL_URL
)

func main() {

	logrus.Info("Configuring.. 🚜")

	//Get routes
	str := util.GetEnv(EnvKeyRoutes, DefaultRoutes)
	logrus.Infoln("..", EnvKeyRoutes, ":", str)
	rules := model.ReadRules(str)

	//Setup routes
	listenAddr := util.GetEnv(EnvKeyListenAddr, DefaultListenAddress)
	logrus.Infoln("..", EnvKeyListenAddr, ":", listenAddr)
	srv := server.SetupServer(listenAddr, rules)

	logrus.Infoln("Configuration complete. Listening now 💥")

	//Serve
	logrus.Fatalln(srv.ListenAndServe())

}
