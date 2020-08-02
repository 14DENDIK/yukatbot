package main

import (
	"flag"
	"log"

	"github.com/14DENDIK/yukatbot/internal/yukat"
	"github.com/14DENDIK/yukatbot/internal/yukat/config"
	"github.com/BurntSushi/toml"
	_ "golang.org/x/text/message/catalog"
)

//go:generate gotext -srclang=en update -out=catalog.go -lang=en,ru,uz

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/yukat.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := config.New()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	server, err := yukat.New(config)
	if err != nil {
		log.Fatal(err)
	}
	if err = server.Start(); err != nil {
		log.Fatal(err)
	}
}
