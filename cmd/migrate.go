package cmd

import (
	"fmt"
	"log"

	"github.com/mojtabafarzaneh/blog/internal/modules/models"
	"github.com/mojtabafarzaneh/blog/pkg/database"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(MigrateCmd)
}

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate",
	Long:  "migrate",
	Run: func(cmd *cobra.Command, args []string) {
		Migrate()
	},
}

func Migrate() {

	db := database.Connect()

	if err := db.AutoMigrate(&models.User{}, &models.Article{}); err != nil {
		log.Fatal("couldn't done the migration", err)
	}

	fmt.Println("migration has been done.")

}
