package cmd

import (
	"github.com/theproshin/images/build"
	"github.com/spf13/cobra"
)

var (
	operaCmd = &cobra.Command{
		Use:   "opera",
		Short: "build Opera image",
		RunE: func(cmd *cobra.Command, args []string) error {
			req := build.Requirements{
				BrowserSource:  build.BrowserSource(browserSource),
				BrowserChannel: browserChannel,
				DriverVersion:  driverVersion,
				NoCache:        noCache,
				TestsDir:       testsDir,
				RunTests:       test,
				IgnoreTests:    ignoreTests,
				Tags:           tags,
				PushImage:      push,
			}
			opera := &build.Opera{req}
			return opera.Build()
		},
	}
)
