package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() (cmd *cobra.Command) {
	var rootCmd = &cobra.Command{
		Use:   "go-cli-template",
		Short: "A simple Go CLI template.",
	}

	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "Prints a friendly greeting",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name == "" {
				fmt.Println("Hello, World!")
			} else {
				fmt.Printf("Hello, %s!\n", name)
			}
		},
	}

	helloCmd.Flags().StringP("name", "n", "", "Specify a name")

	rootCmd.AddCommand(helloCmd)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
