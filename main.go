package main

import (
	"fmt"
	"os"
	"sync"

	gologger "github.com/wawakakakyakya/GolangLogger"
	"github.com/wawakakakyakya/check_logs_by_mail/config"
)

var (
	wg sync.WaitGroup
)

func main() {

	config, err := config.NewConfig()
	if err != nil {
		fmt.Println("[ERROR] read config failed")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	logger := gologger.NewLogger("main", config.LogConfig.Path, config.LogConfig.MaxSize, config.LogConfig.MaxBackups, config.LogConfig.MaxAge, config.LogConfig.Compress, config.LogConfig.Level)
	logger.Info("aaa")
	errCh := make(chan error)
	defer close(errCh)

	logger.Info("done.")
}
