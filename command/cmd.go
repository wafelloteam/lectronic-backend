package command

import (
	"github.com/spf13/cobra"
	"github.com/wafellofazztrack/lectronic-backend/database/orm"
	"github.com/wafellofazztrack/lectronic-backend/lib/server"
)

var initCommand = cobra.Command{
	Short: "lextronic backend",
}

func init() {
	initCommand.AddCommand(server.ServeCmd)
	initCommand.AddCommand(orm.MigrateCmd)
	initCommand.AddCommand(orm.SeedCmd)

}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
