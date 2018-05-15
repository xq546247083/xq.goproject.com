// +build !windows

package dup

import (
	"os"
	"syscall"
)

func init() {
	SetStderr()
}

func SetStderr() {
	// 标准输出
	stdoutFile, _ := os.OpenFile("LOG/Stdout.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0644)
	syscall.Dup2(int(stdoutFile.Fd()), 1)

	// 标准错误输出
	stderrFile, _ := os.OpenFile("LOG/Stderr.txt", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	syscall.Dup2(int(stderrFile.Fd()), 2)
}
