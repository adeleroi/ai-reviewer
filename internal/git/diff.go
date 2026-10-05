package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetStateDiff(pathspec string) (string, error) {
	cmd := exec.Command("git", "diff", pathspec)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git diff failed (%w): %s", err, stderrBuf.String())
	}
	return stdoutBuf.String(), nil
}
