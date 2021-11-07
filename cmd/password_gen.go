package cmd

import (
	"crypto/rand"
	"github.com/spf13/cobra"
)

type Options struct {
	// 桁数指定オプション
	digit int
	// 英小文字オプション
	char bool
	// 英大文字オプション
	CHAR bool
	// 記号オプション
	symbol bool
}

var (
	o = &Options{}
)

func GeneratePasswordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gp",
		Short: "Generate password.",
		Long:  `Generate password command.`,
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("show called: -d: %v\n", o.d)
			n, err := generatePassword(*o)

			if err != nil {
				panic(err)
			}
			cmd.Printf(n)
		},
	}
	cmd.Flags().IntVarP(&o.digit, "digit", "d", 8, "Set the number of digits in the password. Default number of digits is 8.")
	cmd.Flags().BoolVarP(&o.char, "char", "c", false, "Include lowercase letters in the generated password.")

	return cmd
}

func generatePassword(options Options) (string, error) {
	var baseChars = "1234567890"

	// 英小文字オプションあり
	if options.char {
		baseChars += "abcdefghijklmnopqrstuvwxyz"
	}

	buffer := make([]byte, options.digit)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	var result string
	for _, v := range buffer {
		result += string(baseChars[int(v)%len(baseChars)])
	}

	return result, nil
}

func init() {

}
