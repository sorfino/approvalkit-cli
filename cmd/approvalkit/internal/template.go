package internal

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sorfino/approvalkit-cli/internal/git"
)

var (
	_excluded = []string{".fury", "Dockerfile.runtime", ".git", "CHANGELOG.md"}
)

func CopyFromTemplate() {
	scaffoldingDir, err := git.Clone("git@github.com:mercadolibre/fury_mp-approval-go-prj-template.git")
	if err != nil {
		fmt.Printf("fatal :%v", err)
		os.Exit(-1)
	}

	defer os.RemoveAll(scaffoldingDir)

	destinationDir, _ := os.Getwd()
	if err := CleanUp(destinationDir); err != nil {
		fmt.Printf("fatal :%v", err)
		os.Exit(-1)
	}

	if err := git.CheckoutIndex(scaffoldingDir, destinationDir); err != nil {
		fmt.Printf("error :%v.\n", err)
	}
}

func Contains(a []string, x string) bool {
	for i := range a {
		if x == a[i] {
			return true
		}
	}
	return false
}

func CleanUp(dirName string) error {
	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		return err
	}

	for i := range entries {
		if Contains(_excluded, entries[i].Name()) {
			continue
		}

		if err := os.RemoveAll(entries[i].Name()); err != nil {
			return err
		}
	}

	return nil
}
