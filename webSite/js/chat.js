;
//聊天服务器客户端
var webSocketClient = null;

var ChatMain = {
    // 连接
    Connect: function() {
        return connect.call(this);
    },

    // 断开连接
    Disconnect: function() {
        return disconnect.call(this);
    },

    // 发送消息
    SendMessage: function(method, message) {
        return sendMessage.call(this, method, message);
    },
};


// 连接  
function connect() {
    if (webSocketClient != null) {
        return;
    }

    var wsAddress = "ws://" + WebMain.ChatServerConfig + "ws"
    if (window["WebSocket"]) {
        webSocketClient = new WebSocket(wsAddress);
    } else if ('MozWebSocket' in window) {
        webSocketClient = new MozWebSocket(wsAddress);
    } else {
        setStatus("您的浏览器不支持WebSocket");
        return;
    }

    webSocketClient.onopen = function() {
        setStatus("已连接");
    }
    webSocketClient.onmessage = function(e) {
        handerSocketData(e.data)
    }
    webSocketClient.onclose = function(e) {
        setStatus("连接关闭");
    }
    webSocketClient.onerror = function(e) {
        setStatus("连接错误");
    }
}

// 断开连接  
function disconnect() {
    if (webSocketClient != null) {
        webSocketClient.close();
        webSocketClient = null;
    }
}

// 断开连接  
function sendMessage(method, message) {
    var userName = $.cookie("UserName");
    //方法参数
    var requestInfo = {
        UserName: userName,
        Token: $.cookie("Token"),
        Data: new Array(userName, message),
    };

    //调用参数
    var params = {
        MethodName: method,
        RequestInfo: requestInfo
    };

    //获取字符串
    var paramStr = JSON.stringify(params);
    var resultParamStr = $.base64.btoa(paramStr);

    if (webSocketClient != null && webSocketClient.readyState == 1) {
        webSocketClient.send(resultParamStr);
    }
}

// 处理消息
function handerSocketData(data) {
    alert(data)
}

// 设置状态
function setStatus(status) {
    $("#chatStatus").html("聊天服务器状态:" + status);
}

$(document).on("click", ".contactPerson", function() {
    var name = $(this).attr("name");
    var userName = $(this).attr("userName");
    $("#chatHead").html("与" + name + "聊天中");
});

// 获取服务器配置
$(function() {
    ChatMain.Connect();
    var timeStr = new Date().toLocaleDateString();
    $("#chatDate").html(timeStr)
        // ChatMain.SendMessage("SendMessgaeInWorld", "hello world 025game.cn");
});