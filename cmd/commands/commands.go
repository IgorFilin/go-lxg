package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var FilinCmd = &cobra.Command{
	Use:   "filin",
	Short: "filin test cli",
	Long: `test`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	  fmt.Println("Выполнение команды filin")
	},
  }

  var StartCmd = &cobra.Command{
    Use:   "start",
    Short: "Start the application",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Starting the application...")
    },
}
  
  func StartExecute() {
	FilinCmd.AddCommand(StartCmd)
  }