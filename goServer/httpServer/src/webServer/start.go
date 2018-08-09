package webServer

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"crypto/x509"

	"crypto/tls"
	"io/ioutil"
	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/logTool"
)

// 监视开启服务，关闭服务通道
func Monitor(monitorServerChan <-chan bool) {
	for {
		select {
		case flag := <-monitorServerChan:
			if flag {
				startAllServer()
			} else {
				closeAllServer()
			}
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// startServer 开启服务
func startServer(serverPort int) {
	//开启服务
	logTool.LogInfo(fmt.Sprintf("Web服务器监听：%d", serverPort))
	fmt.Println(fmt.Sprintf("Web服务器监听：%d", serverPort))

	// 添加服务
	server := &http.Server{Addr: ":" + strconv.Itoa(serverPort), Handler: new(handle)}
	addServer(server)

	if err := server.ListenAndServe(); err != nil {
		removeServer(server)
		logTool.LogError(err.Error())
	}
}

// startServerTLS 开启TLS服务
func startServerTLS(serverPort int) {
	//开启服务
	logTool.LogInfo(fmt.Sprintf("Web https服务器监听：%d", serverPort))
	fmt.Println(fmt.Sprintf("Web https服务器监听：%d", serverPort))

	// 创建证书池子
	pool := x509.NewCertPool()
	addTrust(pool, configTool.Ca)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(serverPort),
		Handler: new(handle),
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}}

	// 添加服务
	addServer(server)

	if err := server.ListenAndServeTLS(configTool.Crt, configTool.Key); err != nil {
		removeServer(server)
		fmt.Println(err.Error())
		logTool.LogError(err.Error())
	}
}

// 添加信任
func addTrust(pool *x509.CertPool, path string) {
	aCrt, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	pool.AppendCertsFromPEM(aCrt)
}
