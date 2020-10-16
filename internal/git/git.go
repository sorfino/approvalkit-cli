package git

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func Clone(URL string) (string, error) {
	tmpDir := path.Join(os.TempDir(), "fury_go-app-scaffolding")
	if err := os.RemoveAll(tmpDir); err != nil {
		return tmpDir, fmt.Errorf("prepare: %w", err)
	}

	cmd := exec.Command("git", "clone", URL, tmpDir)

	if err := cmd.Start(); err != nil {
		return tmpDir, fmt.Errorf("exec: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return tmpDir, fmt.Errorf("exec wait: %w", err)
	}

	return tmpDir, nil
}

func CheckoutIndex(repositoryPath string, outputDir string) error {
	//git checkout-index -a --prefix=/tmp/scaffoling/

	wd, _ := os.Getwd()

	if outputDir[len(outputDir)-1] != '/' {
		outputDir += "/"
	}

	if err := os.Chdir(repositoryPath); err != nil {
		return fmt.Errorf("chdir: %w", err)
	}

	defer os.Chdir(wd) //nolint

	cmd := exec.Command("git", "checkout-index", "-a", fmt.Sprintf("--prefix=%s", outputDir))
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("exec wait: %w", err)
	}

	return nil
}

func NewPullRequest(repositoryPath, base, body, title string) error {
	return nil
}
