package orm

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/model"
	"github.com/wafellofazztrack/lectronic-backend/database/orm/seeder"
	"gorm.io/gorm"
)

type seederData struct {
	name  string
	model interface{}
	size  int
}

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "db seeder",
	RunE:  Seed,
}

var seedUpVar, seedDownVar bool

func init() {
	SeedCmd.Flags().BoolVarP(&seedUpVar, "up", "u", true, "run seed up")
	SeedCmd.Flags().BoolVarP(&seedDownVar, "down", "d", false, "run seed down")
}

func Seed(cmd *cobra.Command, args []string) error {
	var err error
	db, err := NewDB()
	if err != nil {
		return err
	}

	if seedDownVar {
		err = seedDown(db)
		return err
	}

	if seedUpVar {
		err = seedUp(db)
	}

	return err
}

func seedUp(db *gorm.DB) error {
	var err error

	var seedModel = []seederData{
		{
			name:  "product",
			model: seeder.ProductSeed,
			size:  cap(seeder.ProductSeed),
		},
	}

	for _, data := range seedModel {
		log.Println("create seeding data for", data.name)
		err = db.CreateInBatches(data.model, data.size).Error
	}

	return err

}

func seedDown(db *gorm.DB) error {
	var err error

	var seedModel = []seederData{
		{
			name:  "products",
			model: model.Product{},
		},
	}

	for _, data := range seedModel {
		log.Println("delete seeding data for", data.name)
		sql := fmt.Sprintf("DELETE FROM %v", data.name)
		err = db.Exec(sql).Error
	}

	return err
}
