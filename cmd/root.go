package cmd

import (
"fmt"
"os"
"github.com/joho/godotenv"
"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "up",
}

func init() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Patterns-of-Enterprise-Application-Architecture/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
