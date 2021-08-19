package cmd

import (
	"strings"

	bf "compleat/brainfuck"
	"compleat/converter"
	m "compleat/midipackage"

	"github.com/spf13/cobra"
)

var outputName string
var scaleStr string
var delta uint

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate [input_file] [flags]",
	Short: "Translate from brainfuck to Compleat and vice versa",
	Long: `Translate a brainfuck file to Compleat and a Compleat file to brainfuck.
Neither the translated file nor the original source file will be executed.
Make sure the Comments on your brainfuck files don't contain operators since that could result in some translation errors.

The output file name will follow the following conventions:
	compleat translate input_file_name.bf
	compleat translate input_file_name.b
		--> input_file_name.midi

	compleat translate input_file_name.midi
	compleat translate input_file_name.mid
		--> input_file_name.bf

	compleat translate input_file -o output_file
	compleat translate input_file --output output_file
		--> output_file`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scale, err := m.NewScale(scaleStr)
		if err != nil {
			return err
		}

		file := args[0]
		if outputName == "" {
			outputName = strings.Split(file, ".")[0]
		} else {
			outputName = strings.Split(outputName, ".")[0]
		}

		if strings.HasSuffix(file, ".mid") || strings.HasSuffix(file, ".midi") {
			program, err := converter.MidiToBF(file, scale)
			if err != nil {
				return err
			}

			bf.WriteFile(outputName+".bf", program)

		} else if strings.HasSuffix(file, ".bf") || strings.HasSuffix(file, ".b") {
			program, err := converter.BfToMidi(file, scale)
			if err != nil {
				return err
			}
			err = program.Write(outputName + ".midi", uint32(delta))
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	translateCmd.Flags().StringVarP(&outputName, "output", "o", "", "Name of the output file")
	translateCmd.Flags().StringVarP(&scaleStr, "scale", "s", "CM", "Scale to be used when processing the file")
	translateCmd.Flags().UintVarP(&delta, "delta", "d", 360, "Delta time of every midi note")
}
