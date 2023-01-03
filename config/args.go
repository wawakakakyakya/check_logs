package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/wawakakakyakya/check_logs_by_mail/version"
)

var (
	configPath  string = ""
	showVersion bool   = false
)

func init() {
	flag.StringVar(&configPath, "config", "config.yml", "config path, default ./config.yml")
	flag.BoolVar(&showVersion, "v", false, "show application version")
	flag.BoolVar(&showVersion, "version", false, "show application version")
	flag.Parse()

	if showVersion {
		fmt.Printf("version(%s.%s)\n", version.Version, version.Revision)
		os.Exit(0)
	}
}
