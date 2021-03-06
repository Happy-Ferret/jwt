package main

import (
	"os"

	"github.com/apex/log"
	"github.com/romeovs/jwt/cmd"
	"github.com/romeovs/jwt/util/handler"
)

func main() {
	log.SetHandler(handler.New(os.Stderr))

	err := cmd.RootCmd.Execute()
	if err != nil {
		log.WithError(err).Fatal("Uncaught error")
	}
}
