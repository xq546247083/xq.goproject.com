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
			"-f      toMysql的MD文件路径(相对路径)." + "\n" +
			"-h      host.     默认 127.0.0.1:3306" + "\n" +
			"-u      username. 默认 root" + "\n" +
			"-p      password. 默认 12345678" + "\n" +
			"-d      database. 默认 test" +
			"")
		os.Exit(0)
	}

	flag.StringVar(&core.ToMysqlMdFilePath, "f", "", "toMysql的MD文件路径")
	flag.StringVar(&core.MysqlNetAddress, "h", "127.0.0.1:3306", "数据库地址")
	flag.StringVar(&core.MysqlUserName, "u", "root", "数据库用户名")
	flag.StringVar(&core.MysqlPwd, "p", "12345678", "数据库密码")
	flag.StringVar(&core.DataBaseName, "d", "test", "数据库名称")
	flag.Parse()
}

// 程序入口
func main() {
	if core.ToMysqlMdFilePath == "" {
		core.ToMd()
	} else {
		core.ToMysql()
	}
}
