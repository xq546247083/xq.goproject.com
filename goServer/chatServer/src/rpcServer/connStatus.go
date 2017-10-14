package rpcServer

type connStatus int

const (
	//连接打开
	connOpen = 1 + iota

	//连接等待关闭
	connWaitForClose

	//连接关闭
	connClose
)
