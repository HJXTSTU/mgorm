package boot

import (
	"projects/mgorm/global"
	"os"
	"os/signal"
	"syscall"
)

func init(){
	exit_signal := make(chan os.Signal,1)
	signal.Notify(exit_signal,syscall.SIGINT,syscall.SIGKILL,syscall.SIGTERM)
	global.Mgorm_Init(exit_signal)
}
