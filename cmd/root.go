package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "RNDRUP",
	Short: "automation for RNDR uploads",
	Long:  `an automation time for rndr orbx uploads.`,

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("\n welcome to rndrup cli")

		time.Sleep(2 * time.Second)
	},
}

var send = &cobra.Command{
	Use:   "up",
	Short: "send to rndr",
	Long:  "send to rndr but longer",

	Run: func(cmd *cobra.Command, args []string) {
		Send()
	},
}

// Execute cmd
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
