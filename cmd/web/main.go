package main

import "github.com/victorbischoff/GOFIBER-TMPL/pkg/server"

const port = ":1312"

func main() {
	server.Server(port)

}
