package cmd

import (
	"crypto/rand"
	"errors"
	"fmt"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			password, err := generatePassword(*o)

			if err != nil {
				return err
			}
			cmd.Printf("%v", password)
			fmt.Println()
			return nil
		},
	}
	cmd.Flags().IntVarP(&o.digit, "digit", "d", 8, "Set the number of digits in the password. Default number of digits is 8.")
	cmd.Flags().BoolVarP(&o.char, "char", "c", false, "Include lowercase letters in the generated password.")
	cmd.Flags().BoolVarP(&o.CHAR, "CHAR", "C", false, "Include uppercase letters in the generated password.")
	cmd.Flags().BoolVarP(&o.symbol, "symbol", "s", false, "Include symbols letters in the generated password.")

	return cmd
}

func generatePassword(options Options) (string, error) {
	// 桁指定オプションで7以下のときエラーを返す
	if options.digit < 8 {
		return "", errors.New("the minimum number of digits should be at least 8 characters")
	}
	var baseChars = "1234567890"

	// 英小文字オプションあり
	if options.char {
		baseChars += "abcdefghijklmnopqrstuvwxyz"
	}

	// 英大文字オプションあり
	if options.CHAR {
		baseChars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if options.symbol {
		baseChars += "!@#$%^&*"
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
