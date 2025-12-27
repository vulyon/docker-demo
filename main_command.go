package main

import (
	"fmt"
	"mydocker/container"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/urfave/cli/v2"
)

// run 执行入口
// 判断参数是否包含command
// 调用Run函数执行
var runCommand = cli.Command{
	Name: "run",
	Usage: `reate a container with namespace and cgroups limit
	eg: mydocker run -it [command]`,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "it",
			Usage: "ennable tty",
		},
	},
	Action: func(context *cli.Context) error {
		if context.Args().Len() < 1 {
			return fmt.Errorf("missing container command")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("it")
		Run(cmd, tty)
		return nil
	},
}

// 执行容器初始化操作，RUN() 也会执行init
var initCommand = cli.Command{
	Name:  "init",
	Usage: "init container process run user's process in container.Do not call it outside",
	Action: func(context *cli.Context) error {
		log.Infof("init come on")
		cmd := context.Args().Get(0)
		log.Infof("command %s", cmd)
		err := container.RunContainerInitProcess(cmd, []string{})
		return err
	},
}
