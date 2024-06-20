package core

import (
	"fmt"
	"github.com/zhangga/luban/pkg/logger"
	"strings"
)

type SimpleLauncher struct {
	logger  logger.Logger
	options map[string]string
}

func NewSimpleLauncher(logger logger.Logger) *SimpleLauncher {
	return &SimpleLauncher{
		logger: logger,
	}
}

func (s *SimpleLauncher) Start(xargs ...string) {
	s.options = parseOptions(xargs...)
}

func parseOptions(xargs ...string) map[string]string {
	options := make(map[string]string)
	for _, arg := range xargs {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			err := fmt.Errorf("invalid xargs: %s", arg)
			panic(err)
		}
		if _, ok := options[kv[0]]; ok {
			err := fmt.Errorf("duplicate xargs: %s", arg)
			panic(err)
		}
		options[kv[0]] = kv[1]
	}
	return options
}
