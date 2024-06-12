package cmd

import (
	"github.com/aerokube/images/build"
	"github.com/spf13/cobra"
)

var (
	selenoidVersion string
	seleniumVersion string

	firefoxCmd = &cobra.Command{
		Use:   "firefox",
		Short: "build Firefox image",
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
			firefox := &build.Firefox{SelenoidVersion: selenoidVersion, SeleniumVersion: seleniumVersion, Requirements: req}
			return firefox.Build()
		},
	}
)

func init() {
	firefoxCmd.Flags().StringVar(&selenoidVersion, "selenoid-version", build.LatestVersion, "Selenoid binary version")
	firefoxCmd.Flags().StringVar(&seleniumVersion, "selenium-version", "", "Selenium JAR version")
}
