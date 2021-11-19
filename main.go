package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lmxdawn/vue-admin-go/internal"
	"github.com/lmxdawn/vue-admin-go/model"
	"github.com/lmxdawn/vue-admin-go/router"
	"github.com/urfave/cli"
	"os"
)

// 文档Handle
var swagHandler gin.HandlerFunc

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
func main() {
	app := cli.NewApp()
	app.Name = "vue-admin-go"
	app.Usage = "vue-admin-go -c config/config.yml"
	printVersion := false
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config/config.yml",
			Usage: "config/config.{local|dev|test|pre|prod}.yml",
		},
		cli.BoolFlag{
			Name:        "version, v",
			Required:    false,
			Usage:       "-v",
			Destination: &printVersion,
		},
	}

	app.Action = func(c *cli.Context) error {

		if printVersion {
			fmt.Printf("{%#v}", internal.Get())
			return nil
		}

		conf := c.String("conf")
		err := internal.Init(conf)
		if err != nil {
			return err
		}

		err = model.Database(internal.Config.MySQL)
		if err != nil {
			return err
		}
		server := router.InitRouter()

		if swagHandler != nil {
			server.GET("/swagger/*any", swagHandler)
		}

		server.Run(fmt.Sprintf(":%v", internal.Config.App.Port))

		return nil
	}
	app.Run(os.Args)
}
