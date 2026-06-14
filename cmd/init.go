package cmd

import (
	"os"
	"os/exec"

	"github.com/indium114/jots/internal/storage"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a Git repo inside ~/.jots",
	Run: func(cmd *cobra.Command, args []string) {
		dir := storage.BaseDir()

		// init
		initCommand := exec.Command("git", "-C", dir, "init")
		initCommand.Stdout = os.Stdout
		initCommand.Stderr = os.Stderr
		initCommand.Stdin = os.Stdin
		initCommand.Run()

		// change branch to main
		branchCommand := exec.Command("git", "-C", dir, "branch", "-m", "main")
		branchCommand.Stdout = os.Stdout
		branchCommand.Stderr = os.Stderr
		branchCommand.Stdin = os.Stdin
		branchCommand.Run()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
