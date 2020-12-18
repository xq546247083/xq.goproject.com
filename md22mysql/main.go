package main

import (
	_ "github.com/go-sql-driver/mysql"

	"flag"
	"fmt"
	"os"
	"xq.goproject.com/md22mysql/core"
)

// 初始化参数
func init() {
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage: md22mysql [options...]\n" +
			"--help  This help text" + "\n" +
			"-h      host.     default 127.0.0.1:3306" + "\n" +
			"-u      username. default root" + "\n" +
			"-p      password. default 12345678" + "\n" +
			"-d      database. default test" +
			"")
		os.Exit(0)
	}

	flag.StringVar(&core.MysqlNetAddress, "h", "127.0.0.1:3306", "数据库地址")
	flag.StringVar(&core.MysqlUserName, "u", "root", "数据库用户名")
	flag.StringVar(&core.MysqlPwd, "p", "12345678", "数据库密码")
	flag.StringVar(&core.DataBaseName, "d", "test", "数据库名称")
	flag.Parse()
}

func main() {
	core.ToMd()
}
