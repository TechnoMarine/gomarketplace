package server

import "gomarketplace/config"

func Init() {
	conf := config.GetConfig()
	r := NewRouter()
	r.Run(conf.GetString("server.port"))
}
