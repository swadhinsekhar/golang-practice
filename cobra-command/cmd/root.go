package cmd

import (

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use: "device-commands",
  Short: "Shows supported devices commands",
  Long: `List of registered devices`,
	Example: `- some example of the commands`,

}

func Execute() {
  //on error cobra stack will do os.Exit()
  cobra.CheckErr(rootCmd.Execute())
}
