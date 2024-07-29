package cmd

import (
	"app/connections"
	"app/store"

	"github.com/spf13/cobra"
)

var labCmd = &cobra.Command{
	Use:   "lab",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.SetReportCaller(false)
		db, err := connections.NewSqlite(dbDNS)
		failOnError(err)

		rdb0, err := connections.NewRedis(r0URL)
		failOnError(err)

		cache := connections.NewCache()

		store := store.New(db, rdb0, cache)
		user, err := store.User.GetByUsername("lord0")
		failOnError(err)
		if user == nil {
			log.Warn("User not found")
		} else {
			result, err := store.Room.GetRoomsOfUser(user)
			failOnError(err)
			if result == nil {
				log.Warn("No Rooms")
			} else {
				log.Info(result)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(labCmd)
}
