package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var statusCommand = & cobra.Command {
  Use: "status",
  Short: "Status of daemon",
  Long: `Check the status runinng of daemon`,
  Example: `daemon status`,
  Run: statusRun,
}

func init() {
	rootCmd.AddCommand(statusCommand)
}

func statusRun(cmd *cobra.Command, args []string) {
  logrus.Println("fethcing daemon status")

}
