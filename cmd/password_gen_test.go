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
		cmd.Execute()

		digit := len(buf.String())
		if c.want != digit {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, digit)
		}
	}
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
				if c == baseChar {
					return
				}
			}
		}
		t.Errorf("unexpected response: want:%+v, get:%+v", c.want, buf.String())
	}
}
