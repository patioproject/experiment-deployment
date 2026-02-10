package types

import "os/exec"

func DoCommandType(ctype string, cmd *exec.Cmd) func() ([]byte, error) {
	switch ctype {
	case "CombinedOutput":
		return cmd.CombinedOutput
	default:
		return cmd.Output
	}
}

func CreateCommand(query string) *exec.Cmd {
	return exec.Command("/bin/sh", "-c", query)
}
