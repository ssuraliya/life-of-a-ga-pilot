package cmd

import (
	"context"
	"log"
	"os"

	db "life-of-a-ga-pilot/internal/app/db"

	"github.com/urfave/cli/v3"
)

func Run() {

	dbObject := db.InitDb()

	commandHandler := NewCommandHandler(dbObject)

	cmd := &cli.Command{
		Name:        "Life of a GA Pilot",
		Description: "Application to generate interesting general aviation missions to make flight sim flying interesting!",
		Commands:    commandHandler,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
