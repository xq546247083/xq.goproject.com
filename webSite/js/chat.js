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

    if ('WebSocket' in window) {
        webSocketClient = new WebSocket("ws:" + WebMain.ChatServerConfig);
    } else if ('MozWebSocket' in window) {
        webSocketClient = new MozWebSocket("ws:" + WebMain.ChatServerConfig);
    } else {
        alert("您的浏览器不支持WebSocket。");
        return;
    }

    webSocketClient.onopen = function() {
        setStatus("已连接");
    }
    webSocketClient.onmessage = function(e) {
        handle(e.data)
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
    //方法参数
    var requestInfo = new Array();
    requestInfo["UserName"] = $.cookie("UserName");
    requestInfo["Token"] = $.cookie("Token");
    requestInfo["Data"] = new Array(message)

    //调用参数
    var params = {
        MethodName: method,
        RequestInfo: requestInfo
    };

    //获取字符串并加密
    var paramStr = JSON.stringify(params);
    if (webSocketClient != null) {
        webSocketClient.send(paramStr);
    }
}

// 处理消息
function hander(data) {
    aler(data)
}

// 设置状态
function setStatus(status) {
    $("#chatStatus").html("聊天服务器状态:" + status);
}

// 获取服务器配置
$(function() {
    ChatMain.Connect();

    setTimeout(function() {
        ChatMain.SendMessage("/API/Chat/SendMessgaeInWorld", "hello world");
    }, 1000);
});