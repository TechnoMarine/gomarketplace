package main

import (
	"flag"
	"fmt"
	"github.com/vsouza/go-gin-boilerplate/server"
	"gomarketplace/config"
	//"gomarketplace/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
