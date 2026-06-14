package cmd

import (
	"os"
	"os/exec"

	"github.com/indium114/jots/internal/storage"
	"github.com/spf13/cobra"
)

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Set the Git remote to push your jots to",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		remove := exec.Command("git", "-C", storage.BaseDir(), "remote", "remove", "origin")
		remove.Stdout = os.Stdout
		remove.Stdin = os.Stdin
		remove.Stderr = os.Stderr
		remove.Run()

		command := exec.Command("git", "-C", storage.BaseDir(), "remote", "add", "origin", args[0])
		command.Stdout = os.Stdout
		command.Stdin = os.Stdin
		command.Stderr = os.Stderr
		command.Run()

		show := exec.Command("git", "-C", storage.BaseDir(), "remote", "-v")
		show.Stdout = os.Stdout
		show.Run()
	},
}

func init() {
	rootCmd.AddCommand(remoteCmd)
}
