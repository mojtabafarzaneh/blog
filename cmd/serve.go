package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/blog/pkg/config"
	"github.com/mojtabafarzaneh/blog/pkg/database"
	"github.com/mojtabafarzaneh/blog/pkg/handler"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(ServeCmd)
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve app on the dev server",
	Long:  "application will be served on the port and host provided in config.yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func Serve() {
	config.Set()

	configs := config.Get()

	database.Connect()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})
	c := handler.NewControler()

	router.GET("/articles", c.HandleGetArticles)

	if err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)); err != nil {
		log.Fatal(err)
		return
	}
}
