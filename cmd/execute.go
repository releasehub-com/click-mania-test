package cmd

import (
	"aurora/web"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var start = &cobra.Command{
	Use:   "start",
	Short: "start webserver",
	Long: `Start a webserver that is connected to aurora DB (RDS)
				  that can be used with release to test different scenario`,
	Run: func(cmd *cobra.Command, args []string) {
		web.Setup()
		web.Serve()
	},
}

var setup = &cobra.Command{
	Use:   "setup",
	Short: "setup database",
	Long:  `Create tables on Aurora if it doesn't exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		web.Setup()
	},
}

func Execute() {
	root := &cobra.Command{}
	root.AddCommand(start, setup)
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
