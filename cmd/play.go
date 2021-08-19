package cmd

import (
	"os/exec"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play [input_file]",
	Short: "Play a midi file",
	Long: `Play a midi file.
play uses *timidity*, so you need to have it installed to use this command.
If you use an Ubuntu-based distro you can run the following:
	sudo apt install timidity
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := exec.Command("timidity", args[0]).Run()
		return err
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
