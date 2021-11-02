package cmd

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestGeneratePasswordCmd(t *testing.T) {
	cases := []struct {
		command string
		want    int
	}{
		{command: "cmd-test pg", want: 8},
		//{command: "cmd-test show --int 10", want: "show called: optint: 10, optstr: default"},
		//{command: "cmd-test show --str test", want: "show called: optint: 0, optstr: test"},
	}

	for _, c := range cases {
		buf := new(bytes.Buffer)
		cmd := GeneratePasswordCmd()
		cmd.SetOutput(buf)
		cmdArgs := strings.Split(c.command, " ")
		fmt.Printf("cmdArgs %+v\n", cmdArgs)
		cmd.SetArgs(cmdArgs[1:])
		cmd.Execute()

		digit, _ := strconv.Atoi(buf.String())
		if c.want != countDigits(digit) {
			t.Errorf("unexpected response: want:%+v, get:%+v", c.want, countDigits(digit))
		}
	}
}

func countDigits(n int) (count int) {
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}
