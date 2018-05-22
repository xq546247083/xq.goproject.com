package main

import (
	"fmt"
)

func main() {
	tv := &TV{}
	openCommand := OpenCommand{tv}
	invoker := Invoker{openCommand}
	invoker.Do()

	invoker.SetCommand(CloseCommand{tv})
	invoker.Do()
}

// ----------------执行者----------------

// 执行者
type TV struct{}

func (p *TV) Open() {
	fmt.Println("play...")
}

func (p *TV) Close() {
	fmt.Println("stop...")
}

// ----------------发送者----------------

// 发送者
type Invoker struct {
	cmd Command
}

func (p *Invoker) SetCommand(cmd Command) {
	p.cmd = cmd
}

func (p Invoker) Do() {
	p.cmd.Press()
}

// ------------命令-----------------

// 命令
type Command interface {
	Press()
}

// 打开命令
type OpenCommand struct {
	tv *TV
}

func (p OpenCommand) Press() {
	p.tv.Open()
}

// 关闭命令
type CloseCommand struct {
	tv *TV
}

func (p CloseCommand) Press() {
	p.tv.Close()
}
