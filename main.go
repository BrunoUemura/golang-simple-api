package main

import (
	"github.com/BrunoUemura/golang-simple-api/server"
)

func main() {
	server := server.NewServer()
	server.Run()
}