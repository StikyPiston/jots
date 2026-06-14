package cmd

import (
	"os"
	"os/exec"

	"github.com/indium114/jots/internal/storage"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push your jots to its Git remote",
	Long:  "You can set this remote with `jots remote [url]`",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command("git", "-C", storage.BaseDir(), "push", "origin", "main")
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		command.Run()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
