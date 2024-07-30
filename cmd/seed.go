package cmd

import (
	"fmt"
	"log"

	"github.com/mojtabafarzaneh/blog/internal/modules/models"
	"github.com/mojtabafarzaneh/blog/pkg/database"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rootCmd.AddCommand(SeedCmd)
}

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "seed",
	Long:  "seed",
	Run: func(cmd *cobra.Command, args []string) {
		Seed()
	},
}

func Seed() {
	db := database.Connect()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte("mojtaba7878"), 12)
	if err != nil {
		log.Fatal("couldn't hash the password", err)
		return
	}
	user := models.User{Name: "mojtaba", Email: "mojtaba@gmail.com", Password: string(hashPassword)}

	db.Create(&user)
	fmt.Println("User created")

	for i := 0; i <= 10; i++ {
		article := models.Article{Title: fmt.Sprintf("what shity course %d", i), Content: fmt.Sprintf("what a waste of money %d", i), UserID: 1}
		db.Create(&article)
	}

	fmt.Println("articles are created")

}
