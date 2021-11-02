package cmd

import (
	"crypto/rand"
	"github.com/spf13/cobra"
)

type Options struct {
	// 桁数指定オプション
	d int
	// 英小文字オプション
	c bool
	// 英大文字オプション
	C bool
	// 記号オプション
	s bool
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
			n, err := generatePassword(o.d)

			if err != nil {
				panic(err)
			}
			cmd.Printf(n)
		},
	}
	cmd.Flags().IntVarP(&o.d, "digit", "d", 8, "Set the number of digits in the password. Default number of digits is 8.")
	cmd.Flags().BoolVarP(&o.c, "character", "c", false, "character option")

	return cmd
}

const numberChars = "1234567890"

func generatePassword(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	charsLength := len(numberChars)
	for i := 0; i < length; i++ {
		buffer[i] = numberChars[int(buffer[i])%charsLength]
	}

	return string(buffer), nil
}

func init() {

}
