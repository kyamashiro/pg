package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// 桁数オプションテスト
func TestGeneratePasswordCmd_OptionDigit(t *testing.T) {
	cases := []struct {
		command string
		want    int
	}{
		{command: "cmd-test gp", want: 8},
		{command: "cmd-test gp -d 10", want: 10},
		{command: "cmd-test gp --digit 10", want: 10},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := GeneratePasswordCmd()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		err := cmd.Execute()

		if err != nil {
			Expected := "the minimum number of digits should be at least 8 characters"
			if err.Error() != Expected {
				t.Errorf("Error actual = %v, and Expected = %v.", err, Expected)
			}
		}

		digit := len(buf.String())
		if c.want != digit {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, digit)
		}
	}

	t.Run("when error the number of digits is less than 7", func(t *testing.T) {
		cases := []struct {
			command string
			want    int
		}{
			{command: "cmd-test gp --digit 7", want: 7},
		}

		for _, c := range cases {
			buf := new(bytes.Buffer)
			cmd := GeneratePasswordCmd()
			cmd.SetOutput(buf)
			cmdArgs := strings.Split(c.command, " ")
			fmt.Printf("cmdArgs %+v\n", cmdArgs)
			cmd.SetArgs(cmdArgs[1:])
			err := cmd.Execute()

			if err != nil {
				Expected := "the minimum number of digits should be at least 8 characters"
				if err.Error() != Expected {
					t.Errorf("Error actual = %v, and Expected = %v.", err, Expected)
				}
			}
		}
	})
}

// 英小文字オプションテスト
func TestGeneratePasswordCmd_OptionChar(t *testing.T) {
	cases := []struct {
		command string
		want    bool
	}{
		{command: "cmd-test gp -c", want: true},
		{command: "cmd-test gp --char", want: true},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := GeneratePasswordCmd()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		var baseStr = "abcdefghijklmnopqrstuvwxyz"

		for _, c := range buf.String() {
			for _, baseChar := range baseStr {
				// 英子文字が生成されたパスワードに含まれている場合は終了
				if c == baseChar {
					return
				}
			}
		}
		t.Errorf("unexpected response: want:%+v, get:%+v", c.want, buf.String())
	}
}
