package console

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "start the serveer",
	Long: "start the server",
	Run: server,

}

func init() {
	RootCMD.AddCommand(serverCmd);
}

func server(cmd *cobra.Command, args []string) {
	setupLogger()
	fmt.Println("ok")
}