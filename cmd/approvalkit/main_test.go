package main_test

import (
	"testing"

	"github.com/docopt/docopt-go"
	"github.com/stretchr/testify/require"
)

func TestParsingArgs(t *testing.T) {
	usage := `Usage:
 approvalkit  <cmd>
 approvalkit -h | --help

Options:
 -h --help    Show this screen`

	p, err := docopt.ParseDoc(usage)
	require.NoError(t, err)

	require.NotEmpty(t, p)
}
