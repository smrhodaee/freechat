package cmd

import (
	"app/api"
	"app/connections"
	"app/models"
	"app/router"
	"app/services/ratelimit"
	"app/store"
	"app/ws"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := connections.NewMySql(dbDNS)
		failOnError(err)
		failOnError(db.AutoMigrate(
			&models.User{},
			&models.Room{},
			&models.RoomMember{},
			&models.Message{},
			&models.File{},
		))
		rdb0, err := connections.NewRedis(r0URL)
		failOnError(err)
		rdb1, err := connections.NewRedis(r1URL)
		failOnError(err)
		cache := connections.NewCache()
		store := store.New(db, rdb0, cache)
		ratelimit := ratelimit.New(rdb1)
		app := router.New(log, store)
		ws := ws.NewWS(log, store)
		go ws.HandleChannels()
		ws.RegisterRoutes(app)
		api := api.NewAPI(log, store, ratelimit, ws)
		api.Register(app)
		failOnError(app.Listen(addr))
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
