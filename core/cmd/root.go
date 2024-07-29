package cmd

import (
	"app/tools"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var log = tools.NewLogger()

const (
	addr  = ":8080"
	r0URL = "redis://localhost:6379/0"
	r1URL = "redis://localhost:6379/1"
	dbDNS = "lord:123456@tcp(127.0.0.1:3306)/general?parseTime=True"
)

var rootCmd = &cobra.Command{}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
