package orm

import (
	"github.com/spf13/cobra"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	RunE:  migrate,
}

var migrationUp bool
var migrationDown bool

func init() {
	MigrateCmd.Flags().BoolVarP(&migrationUp, "up", "u", true, "running auto migration")
	MigrateCmd.Flags().BoolVarP(&migrationDown, "down", "d", false, "running auto reset migration")
}

func migrate(cmd *cobra.Command, args []string) error {
	db, err := NewDB()
	if err != nil {
		return err
	}

	if migrationDown {
		return db.Migrator().DropTable(&model.User{}, &model.Product{}, &model.History{}, &model.Cart{})
	}

	if migrationUp {
		return db.AutoMigrate(&model.User{}, &model.Product{}, &model.History{}, &model.Cart{})
	}

	return nil

}
