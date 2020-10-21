package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/sorfino/approvalkit-cli/cmd/approvalkit/internal"
)

func main() {
	usage := `Usage:
 approvalkit <cmd>
 approvalkit -h | --help

Options:
 -h --help    Show this screen`

	cmds, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Printf("fatal :%v", err)
		os.Exit(-1)
	}

	cmd, err := cmds.String("<cmd>")
	if err != nil {
		fmt.Printf("fatal :%v", err)
		os.Exit(-1)
	}

	switch {
	case strings.EqualFold(cmd, "version"):
		fmt.Printf("version %s\n", _version)
		os.Exit(0)
	case strings.EqualFold(cmd, "template"):
		fmt.Println("working ...")
		internal.CopyFromTemplate()
		fmt.Println("done.")
	default:
		os.Exit(1)
	}

}
