package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

type CommandHandler struct {
	db *gorm.DB
}

func NewCommandHandler(db *gorm.DB) []*cli.Command {

	commandHandler := CommandHandler{db: db}

	return []*cli.Command{
		{
			Name:   "generate",
			Usage:  "Generate a mission",
			Action: commandHandler.GenerateMission,
		},
		{
			Name:  "player",
			Usage: "Manage players",
			Commands: []*cli.Command{
				{
					Name:   "create",
					Usage:  "Create a new player",
					Action: commandHandler.AddAircraftCommand,
				},
			},
		},
		{
			Name:  "fleet",
			Usage: "Manage fleet",
			Commands: []*cli.Command{
				{
					Name:   "add",
					Usage:  "Add aircraft to collection",
					Action: commandHandler.AddAircraftCommand,
				},
			},
		},
	}
}

func (c *CommandHandler) GenerateMission(ctx context.Context, cmd *cli.Command) error {
	return nil
}

func (c *CommandHandler) AddAircraftCommand(ctx context.Context, cmd *cli.Command) error {
	return nil
}

func (f *CommandHandler) CreatePlayerCommand(ctx context.Context, cmd *cli.Command) error {
	return nil
}
