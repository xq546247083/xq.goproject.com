package main

import (
	"xq.goproject.com/goServerModel/src/rpcServerObject"
	"xq.goproject.com/goServerModel/src/webServerObject"
	"xq.goproject.com/commonTool/intTool"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	testWebServer()

	conn, err := net.DialTimeout("tcp", "192.168.0.102:8187", 2*time.Second)
	if err != nil {
		fmt.Printf("Dial Error: %s", err)
	} else {
		fmt.Printf("Connect to the server. (local address: %s)", conn.LocalAddr())
	}

	var requestObj rpcServerObject.RequestObject
	requestObj.MethodName = "PlayerLogin"
	requestObj.Parameters = []interface{}{"xxx"}

	message, _ := json.Marshal(&requestObj)

	str := append(intTool.Int32ToByte(int32(len(message)), binary.LittleEndian), message...)

	defer func() {
		conn.Close()
	}()

	go func() {
		for {
			conn.Write(str)
			time.Sleep(time.Second * 1000)
		}
	}()

	//获取服务端返回数据
	for {
		readData := make([]byte, 1024)

		//阻塞读取数据
		len, err := conn.Read(readData)
		if err != nil {
			fmt.Printf("客户端等待数据错误")
			break
		}

		fmt.Println(string(readData[:len]))
	}
}

func testWebServer() {
	var requestObj webServerObject.RequestObject
	requestObj.MethodName = "PlayerLogin"
	requestObj.Parameters = []interface{}{"xx3x"}
	requestByte, _ := json.Marshal(requestObj)
	requestStr := string(requestByte)

	response, err := http.Post("http://192.168.0.102:8186/API", "application/x-www-form-urlencoded", strings.NewReader(requestStr))
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
