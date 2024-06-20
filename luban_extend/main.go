package main

import (
	"github.com/zhangga/luban/cmd/luban/rootcmd"
	"log"
)

func main() {
	if err := rootcmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
