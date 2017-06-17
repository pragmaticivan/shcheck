package sh

import (
	"fmt"
	"os/exec"
	"runtime"
)

const shellcheckPath = "/tmp/shellcheck"

type shellcheck struct {
}

// Check a file with shellcheck
func (s *shellcheck) Check(file string) error {
	bin, err := binaryFor(s, "shellcheck")
	if err != nil {
		return err
	}
	out, err := exec.Command(bin, "-x", file).CombinedOutput()
	if err == nil {
		return nil
	}
	return fmt.Errorf("shellcheck failed: %v", string(out))
}

// Install shellcheck
func (*shellcheck) Install() (string, error) {
	if runtime.GOOS == "linux" {
		return shellcheckPath, download(
			"https://github.com/caarlos0/shellcheck-docker/releases/download/v0.4.6/shellcheck",
			shellcheckPath,
		)
	}
	return "", fmt.Errorf("platform not supported: %v", runtime.GOOS)
}
