package main

import (
	"portsApi/internal/config"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	_ = config.Read()
	return nil
}