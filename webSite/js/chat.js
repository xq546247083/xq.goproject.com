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
        //连接成功，广播到所有用户
        ChatMain.SendMessage("BroadClients", "");
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
function handerSocketData(returnData) {
    var userName = $.cookie("UserName");

    var returnObj = JSON.parse(returnData);
    if (returnObj.Type == "BroadClients") {
        var personStr = "  <li><a href=\"#\" class=\"contactPerson\" fullName=\"所有人\" userName=\"所有人\"><i class=\"fa fa-comments \"></i> 所有人</a></li>"
        $.each(returnObj.Data, function(n, value) {
            //如果不是当前用户广播，则添加进入列表
            if (value.UserName != userName) {
                personStr += "<li><a href=\"#\" class=\"contactPerson\" fullName=\"" + value.FullName + "\" userName=\"" + value.UserName + "\"> <i class=\"fa fa-comment\"></i> " + value.FullName + "</a></li>";
            }
        });

        $("#chatPersonList").html(personStr)
    } else if (returnObj.Type == "World") {
        if (userName == returnObj.Data.FromSysUserName) {
            var crTimeStr = returnObj.Data.Crtime.substr(11, 5);
            var messageContent = "<div class=\"right\"><div class=\"author-name\">" + returnObj.Data.FromSysUserName + "<small class=\"chat-date \">" + crTimeStr + "</small></div><div class=\"chat-message active \">" + returnObj.Data.Message + "</div></div>"
            $("#chatContent").append(messageContent)
        } else {
            var crTimeStr = returnObj.Data.Crtime.substr(11, 5);
            var messageContent = "<div class=\"left\"><div class=\"author-name\">" + returnObj.Data.FromSysUserName + "<small class=\"chat-date \">" + crTimeStr + "</small></div><div class=\"chat-message \">" + returnObj.Data.Message + "</div></div>"
            $("#chatContent").append(messageContent)
        }

    } else if (returnObj.Type == "Private") {

    }
}

// 设置状态
function setStatus(status) {
    $("#chatStatus").html("聊天服务器状态:" + status);
}

$(document).on("click", ".contactPerson", function() {
    var name = $(this).attr("fullName");
    var userName = $(this).attr("userName");
    $("#chatHead").html("与 " + name + " 聊天中");
    $("#chatHead").attr("userName") = userName;
});

$(document).on("click", "#sendMessageBtn", function() {
    var talkUserName = $("#chatHead").attr("userName");
    var messgae = $("#messageInput").val()

    if (talkUserName == "所有人") {
        ChatMain.SendMessage("SendMessgaeInWorld", messgae);
    } else {
        //这里是私聊
    }
});


// 获取服务器配置
$(function() {
    ChatMain.Connect();
    var timeStr = new Date().toLocaleDateString();
    $("#chatDate").html(timeStr)

    // setTimeout(function() {
    //     ChatMain.SendMessage("SendMessgaeInWorld", "hello world 025game.cn");
    // }, 1000);
});