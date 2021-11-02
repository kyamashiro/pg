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
		Use:   "pg",
		Short: "Generate password.",
		Long:  `Generate password command.`,
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Printf("show called: optint: %d, optstr: %d", o.d, o.c)
			n, err := generatePassword(8)

			if err != nil {
				panic(err)
			}
			cmd.Printf(n)
		},
	}
	cmd.Flags().IntVarP(&o.d, "digit", "d", 0, "Set the number of digits in the password.")
	cmd.Flags().BoolVarP(&o.c, "character", "c", false, "character option")

	return cmd
}

const otpChars = "1234567890"

func generatePassword(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func init() {

}
