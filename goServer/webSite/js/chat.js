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
    SendMessage: function(method, talkToUserName, message) {
        return sendMessage.call(this, method, talkToUserName, message);
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
        ChatMain.SendMessage("Login", "", "");
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
function sendMessage(method, talkToUserName, message) {
    var userName = $.cookie("UserName");
    //方法参数
    var requestInfo = {
        UserName: userName,
        Token: $.cookie("Token"),
        Data: new Array(userName, talkToUserName, message),
    };

    //调用参数
    var params = {
        MethodName: method,
        RequestInfo: requestInfo
    };

    //获取字符串
    var paramStr = JSON.stringify(params);

    if (webSocketClient != null && webSocketClient.readyState == 1) {
        webSocketClient.send(paramStr);
    }
}

// 处理消息
function handerSocketData(returnData) {
    var userName = $.cookie("UserName");

    var returnObj = JSON.parse(returnData);
    if (returnObj.Type == "BroadClients") {
        var chatPerSonUserName = $("#chatHead").attr("username");

        //遍历返回的元素，如果不存在，则添加
        $.each(returnObj.Data, function(n, value) {
            var existFlag = false;
            var tempNode;
            $("#chatPersonList li").each(function() {
                var liUserName = $($(this).children("a").get(0)).attr("username");
                if (liUserName == value.UserName) {
                    existFlag = true;
                    tempNode = this;
                    return;
                }
            });

            //如果不是当前用户广播，则添加进列表
            if (value.UserName != userName) {
                var chatTargetStatusStr = chatPerSonUserName == value.UserName ? "class=\"chat-back\"" : "";
                var chatStatusStr = value.Online ? "chat-online" : "chat-notonline ";
                var contentLiStr = "<li " + chatTargetStatusStr + "><a href=\"#\" class=\"contactPerson\" fullname=\"" + value.FullName + "\" username=\"" + value.UserName + "\"> <i class=\"fa fa-comment\"></i> <font class=\"" + chatStatusStr + "\">" + value.FullName + "</font></a></li>";

                if (existFlag) {
                    $(tempNode).remove();
                }

                if (value.Online) {
                    $($("#chatPersonList").children("li").get(0)).after(contentLiStr);
                } else {
                    $("#chatPersonList").append(contentLiStr);
                }
            }
        });

        //如果服务器不存在，则删除
        $("#chatPersonList li").each(function() {
            var liUserName = $($(this).children("a").get(0)).attr("username");
            var existFlag = false;
            $.each(returnObj.Data, function(n, value) {
                if (liUserName == value.UserName) {
                    existFlag = true;
                    return;
                }
            });
            if (!existFlag && liUserName != "所有人") {
                $(this).remove();
            }
        });
    } else if (returnObj.Type == "World") {
        var chatPerSonUserName = $("#chatHead").attr("username");

        //获取玩家聊天历史记录
        var history = WebMain.GetLocalData("AllPersonHistory");
        if (history == null || history == "undefined") {
            history = "";
        }

        var crTimeStr = returnObj.Data.Crtime.substr(11, 5);
        var messageContent = "";

        if (userName == returnObj.Data.FromSysUserName) {
            messageContent = "<div class=\"right\"><div class=\"author-name\">" + returnObj.Data.FromSysUserName + "<small class=\"chat-date \">" + crTimeStr + "</small></div><div class=\"chat-message active \">" + returnObj.Data.Message + "</div></div>";

        } else {
            toastr.success("提示", "收到一条广播消息");
            AddMessgaeNum();
            messageContent = "<div class=\"left\"><div class=\"author-name\">" + returnObj.Data.FromSysUserName + "<small class=\"chat-date \">" + crTimeStr + "</small></div><div class=\"chat-message \">" + returnObj.Data.Message + "</div></div>";
        }

        //聊天追加元素，更新聊天记录cookie
        if (chatPerSonUserName == "所有人") {
            $("#chatContent").append(messageContent);
            scrollToEnd();
        }

        WebMain.SaveLocalData("AllPersonHistory", history + messageContent);

    } else if (returnObj.Type == "Private") {
        var chatPerSonUserName = $("#chatHead").attr("username");


        var crTimeStr = returnObj.Data.Crtime.substr(11, 5);
        var messageContent = "";

        if (userName == returnObj.Data.FromSysUserName) {
            //获取玩家聊天历史记录
            var history = WebMain.GetLocalData(returnObj.Data.SysUserName + "History");
            if (history == null || history == "undefined") {
                history = "";
            }

            messageContent = "<div class=\"right\"><div class=\"author-name\">" + returnObj.Data.FromSysUserName + "<small class=\"chat-date \">" + crTimeStr + "</small></div><div class=\"chat-message active \">" + returnObj.Data.Message + "</div></div>";

            //聊天追加元素
            if (chatPerSonUserName == returnObj.Data.SysUserName) {
                $("#chatContent").append(messageContent);
                scrollToEnd();
            }
            WebMain.SaveLocalData(returnObj.Data.SysUserName + "History", history + messageContent);
        } else {
            //获取玩家聊天历史记录
            var history = WebMain.GetLocalData(returnObj.Data.FromSysUserName + "History");
            if (history == null || history == "undefined") {
                history = "";
            }

            messageContent = "<div class=\"left\"><div class=\"author-name\">" + returnObj.Data.FromSysUserName + "<small class=\"chat-date \">" + crTimeStr + "</small></div><div class=\"chat-message \">" + returnObj.Data.Message + "</div></div>";

            //聊天追加元素
            if (chatPerSonUserName == returnObj.Data.FromSysUserName) {
                $("#chatContent").append(messageContent);
                scrollToEnd();
            }

            toastr.success("提示", "收到来自" + returnObj.Data.FromSysUserName + "的私聊消息");
            AddMessgaeNum();
            WebMain.SaveLocalData(returnObj.Data.FromSysUserName + "History", history + messageContent);
        }
    }
}

//点击发送消息
function sendMessageClick() {
    var talkToUserName = $("#chatHead").attr("username");
    var messgae = $("#messageInput").val()
    if (messgae == "") {
        return;
    }

    var method = "BroadMessgae"
    if (talkToUserName != "所有人") {
        method = "SendMessgae"
    }

    ChatMain.SendMessage(method, talkToUserName, messgae);

    $("#messageInput")[0].focus()
    $("#messageInput").val("");
    scrollToEnd();
}


// 设置状态
function setStatus(status) {
    $("#chatStatus").html("聊天服务器状态:" + status);
}

//添加消息数量
function AddMessgaeNum() {
    var htmlStr = $("#messgaeNum").html();
    if (htmlStr == null || htmlStr == "") {
        $("#messgaeNum").html("1");
    } else {
        $("#messgaeNum").html(parseInt(htmlStr) + 1);
    }
}

// 聊天框滚动到底部
function scrollToEnd() {
    var beforeHeight = $("#chatContent").scrollTop();
    $("#chatContent").scrollTop($("#chatContent").scrollTop() + 20);
    var afterHeight = $("#chatContent").scrollTop();
    if (beforeHeight == afterHeight) {

    } else {
        setTimeout("scrollToEnd()", 5);
    }
}

$(document).on("click", ".contactPerson", function() {
    var fullName = $(this).attr("fullName");
    var userName = $(this).attr("username");
    $("#chatHead").html("与 " + fullName + " 聊天中");
    $("#chatHead").attr("username", userName);
    $("#chatHead").attr("fullname", fullName);

    //设置选中行背景
    $("#chatPersonList li").each(function() {
        $(this).removeClass("chat-back");
    });

    $(this).parent().addClass("chat-back");

    //加载聊天记录
    var history = "";
    if (userName == "所有人") {
        //获取玩家聊天历史记录
        history = WebMain.GetLocalData("AllPersonHistory");
    } else {
        history = WebMain.GetLocalData(userName + "History");
    }

    if (history == null || history == "undefined") {
        history = "";
    }

    $("#chatContent").html(history)
    scrollToEnd();
});

$("#chatContent").scroll(function() {
    $("#messgaeNum").html("");
});

$(document).on("click", "#sendMessageBtn", function() {
    sendMessageClick();
});


// 获取服务器配置
$(function() {
    ChatMain.Connect();
    var timeStr = new Date().toLocaleDateString();
    $("#chatDate").html(timeStr)

    //加载所有人的聊天记录
    $("#chatPersonList li:first").addClass("chat-back");
    var history = WebMain.GetLocalData("AllPersonHistory");
    $("#chatContent").html(history)
});

//点击打开消息按钮
$(".open-small-chat").click(function() {
    $("#messgaeNum").html("");
})

//回车提交
$(function() {
    $(document).keydown(function(e) {
        if (e.keyCode == "13") {
            sendMessageClick();
        }
    })
})