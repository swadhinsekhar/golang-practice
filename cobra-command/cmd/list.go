package cmd

import (

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type listCmd struct {
	cobra.Command

	jsonFlag     bool
	propsFlag    bool
	printIDsFlag bool
	stateFlag    string
	userFlag     string
}

func init() {
	var c listCmd
	c.Command = cobra.Command{
		Use:   "list [filters...]",
		Short: "list your devices",
		Long: `list your devices, optionally filtering the results`,
		Run: c.run,
	}
	rootCmd.AddCommand(&c.Command)
}

func (c *listCmd) run(cmd *cobra.Command, filters []string) {
  logrus.Println("list commands executes success")
}
