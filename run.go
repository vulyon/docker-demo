package main

import (
	"os"

	"mydocker/container"

	"github.com/rs/zerolog/log"
)

// start() 才是真正执行命令创建好的command的调用，它首先会clone一个namespace隔离的进程。在子进程中调用/proc/self/exe (也就是自身)，发送init参数，调用我们写的init方法
// 初始化容器的一些资源
func Run(cmd string, tty bool) {
	parent := container.NewParentProcess(cmd, tty)
	if err := parent.Start(); err != nil {
		log.Error().Err(err).Msg("Failed to start parent process")

	}
	_ = parent.Wait()
	os.Exit(-1)
}
