package main

import (
	"flag"
	"os"

	"pecker/commands"

	"github.com/sirupsen/logrus"
)

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		logrus.Info(`
			支持命令：
			init
			model {model_name}
		`)
		return
	}

	if args[0] == "init" {
		err := commands.InitProject()
		if err != nil {
			logrus.Error(err.Error())
		}
		return
	}

	if args[0] == "model" {
		if len(args) < 2 {
			logrus.Error("缺少 model_name")
			os.Exit(2)
		}
		err := commands.AddModel(args[1])
		if err != nil {
			logrus.Error(err.Error())
		}
		return
	}
}
