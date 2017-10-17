package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"xq.goproject.com/commonTools/intTool"
	"xq.goproject.com/goServer/goServerModel/src/rpcServerObject"
)

func main() {
	//testWebServer()
	testRPCServer()
}

func testWebServer() {
	requestObj := make(map[string]interface{})
	data := []interface{}{"xiaoqiang"}
	requestObj["Data"] = data
	requestObj["UserName"] = "xiaoqiang"
	requestObj["Token"] = "5b72ba0ec39982a4d45e11d0236903a3"

	requestByte, _ := json.Marshal(requestObj)
	requestStr := string(requestByte)

	response, err := http.Post("http://localhost:8883/API/SysMenu/GetInfo", "application/x-www-form-urlencoded", strings.NewReader(requestStr))
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

func testRPCServer() {
	conn, err := net.DialTimeout("tcp", "localhost:8886", 2*time.Second)
	if err != nil {
		fmt.Printf("Dial Error: %s", err)
	} else {
		fmt.Printf("Connect to the server. (local address: %s)", conn.LocalAddr())
	}

	requestObj := rpcServerObject.NewRequestObject()
	requestObj.MethodName = "RpcTest"
	requestObj.RequestInfo["Data"] = []interface{}{"xx111x"}

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
