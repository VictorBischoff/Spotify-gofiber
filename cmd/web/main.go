package main

import "github.com/victorbischoff/GOFIBER-TMPL/pkg/server"

const port = ":8080"

func main() {
	server.Server(port)

}
