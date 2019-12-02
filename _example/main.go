package main

import (
	"github.com/miRemid/hanabi"
	"github.com/miRemid/hanabi/plugins"
)

func main() {
	client := hanabi.NewServer("127.0.0.1", 5700)
	client.AccessToken("asdf")
	client.Register(plugins.Roll{
		Cmd:"roll",
		Area: 1000,
	})
	client.Run(":3000", "/")
}
