package webServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"xq.goproject.com/commonTools/configTool"
	"xq.goproject.com/commonTools/logTool"
	"xq.goproject.com/commonTools/stringTool"
	"xq.goproject.com/goServer/goServerModel/src/webServerObject"
)

//Handle webserver服务处理
type handle struct{}

//服务监听
func (handleObj *handle) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestObject := webServerObject.NewRequestObject(request)
	responseObj := webServerObject.NewResponseObject()
	logInfo := ""

	// 应对NLB的监控
	if request.RequestURI == "/" || request.RequestURI == "/favicon.ico" {
		return
	}

	//返回前，先返回数据，后写日志
	defer func() {
		data, err := json.Marshal(responseObj)
		if err != nil {
			logInfo = fmt.Sprintf("返回对象-反序列化出错，请求:%s%s错误信息为：%s", request, stringTool.GetNewLine(), err)
			//返回对象反序列化失败，只能返回空数据
			data = []byte("")
		}

		//返回数据
		responseWriter.Header().Add("Access-Control-Allow-Origin", "*")
		responseWriter.Write(data)

		//记录请求和返回数据
		if logInfo != "" {
			logTool.Log(logTool.Error, logInfo)
		} else {
			if valStr, err := requestObject.GetValStr(); err != nil {
				logTool.LogError(fmt.Sprintf("web服务器获取请求数据失败,请求：%s%serr:%s", request, stringTool.GetNewLine(), err))
			} else {
				logTool.LogDebug(fmt.Sprintf("web服务器接受请求,请求地址：%s %s请求数据：%s %s返回数据：%s", requestObject.HTTPRequest.RequestURI, stringTool.GetNewLine(), valStr, stringTool.GetNewLine(), string(data)))
			}
		}
	}()

	//检测refrer
	httpReferer := request.Referer()
	webSiteRefererList := strings.Split(configTool.WebSiteReferer, ",")
	refererFlag := false
	for _, webSiteReferer := range webSiteRefererList {
		if strings.Index(httpReferer, webSiteReferer) == 0 {
			refererFlag = true
		}
	}

	if !refererFlag {
		responseObj.SetResultStatus(webServerObject.DataError)
		return
	}

	//如果调用的用户API，则检测用户请求
	var PwdExpiredTime interface{}
	if strings.Index(request.RequestURI, "/API") == 0 {
		handlerObj, _ := getHandler("/InnerFunc/SysUser/CheckRequest")
		updatePwdExpiredTimeResponseObj := handlerObj.handlerFunc(requestObject)
		if updatePwdExpiredTimeResponseObj.Status != webServerObject.Success {
			responseObj = updatePwdExpiredTimeResponseObj
			return
		}

		PwdExpiredTime = updatePwdExpiredTimeResponseObj.AttachData["PwdExpiredTime"]
	}

	// 根据路径选择不同的处理方法
	handlerObj, exists := getHandler(request.RequestURI)
	if !exists {
		responseObj.SetResultStatus(webServerObject.APINotDefined)
		return
	}

	// 调用方法,并赋值更新后的密码过期时间
	responseObj = handlerObj.handlerFunc(requestObject)
	responseObj.AttachData["PwdExpiredTime"] = PwdExpiredTime
}
