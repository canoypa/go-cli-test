package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "go_cli_test",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("Hello World!")
		},
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
