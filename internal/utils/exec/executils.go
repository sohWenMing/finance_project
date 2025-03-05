package executils

import "os/exec"

func GenExecCommand(dir string, command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	return cmd
}
