package cmd

import (
	"log"
	"os"

	"fudjie.waizly/backend-test/cmd/database"
	"fudjie.waizly/backend-test/cmd/http"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "waizly-backend",
		Short: "Services",
		Long:  "Services",
	}
)

func Execute() {
	//Register command
	rootCmd.AddCommand(http.ServeHTTPCmd(), database.MigrationCommand)

	http.ServeHTTPCmd().Flags().StringP("config", "c", "config/file", "Config dir i.e. config/file")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("Error: \n", err.Error())
		os.Exit(-1)
	}
}
