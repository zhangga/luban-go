package main

import (
	"github.com/spf13/cobra"
	"github.com/zhangga/luban/pkg/logger"
)

func run(cmd *cobra.Command, args []string) {
	logger, err := logger.DefaultLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Flush()

	loadCommandOptions(configPath, &options)
}
