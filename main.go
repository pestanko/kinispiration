package main

import "github.com/pestanko/kinispiration/kinspiration"

func main() {
	config := kinspiration.Config{}
	config.Init()
	app := kinspiration.App{}
	app.Init(&config)
	app.Run(":3000")
}

