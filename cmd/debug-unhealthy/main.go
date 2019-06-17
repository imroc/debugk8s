package main

import (
	"log"
	"os"

	"github.com/imroc/debugk8s/cmd/debug-unhealthy/app"
)

func main() {
	cmd := app.NewUnhealthyCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("exit")
}
