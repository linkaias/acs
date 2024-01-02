package main

import (
	"APICallerStats/bootstrap"
)

func main() {
	app := bootstrap.App()
	// cmd serve
	app.Run()
}
