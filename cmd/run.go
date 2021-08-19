package cmd

import (
	bf "compleat/brainfuck"
	"compleat/converter"
	m "compleat/midipackage"
	"strings"

	"github.com/spf13/cobra"
)

var runOutputName string
var runScaleStr string
var runDelta uint

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [input_file] [flags]",
	Short: "Run a brainfuck or Compleat file.",
	Long: `Run a brainfuck or Compleat file.
Make sure the Comments on your brainfuck files don't contain operators since that could result in some translation errors.

If you provide an Argument to the -o/--output flag, the file will also be translated with the provided name as the name of the output file.
If you don't, the file will only be interpreted.

Providing an Argument to the -s/--scale flag will change the scale that interprets the file.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		scale, err := m.NewScale(runScaleStr)
		if err != nil {
			return err
		}

		file := args[0]

		if strings.HasSuffix(file, ".mid") || strings.HasSuffix(file, ".midi") {
			program, err := converter.MidiToBF(file, scale)
			if err != nil {
				return err
			}

			_, err = bf.RunBF(program)
			if err != nil {
				return err
			}

			if runOutputName != "" {
				bf.WriteFile(runOutputName, program)
			}

		} else if strings.HasSuffix(file, ".bf") || strings.HasSuffix(file, ".b") {
			_, err = bf.RunFile(file)
			if err != nil {
				return err
			}

			if runOutputName != "" {

				program, err := converter.BfToMidi(file, scale)
				if err != nil {
					return err
				}

				err = program.Write(runOutputName, uint32(runDelta))
				if err != nil {
					return err
				}
			}

		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&runOutputName, "output", "o", "", "Name of the output file")
	runCmd.Flags().StringVarP(&runScaleStr, "scale", "s", "CM", "Scale to be used when processing the file")
	runCmd.Flags().UintVarP(&runDelta, "delta", "d", 360, "Delta time of every midi note")
}
