﻿var WebMain = {
    //----------------------------------------------一些配置数据----------------------------------------------
    //业务服务器配置
    WebServerConfig: "",

    //文件服务器配置
    FileServerConfig: "",

    //聊天服务器配置
    ChatServerConfig: "",

    //----------------------------------------------一些通用方法----------------------------------------------
    //初始化,检测数据
    //flag: 0：登录页面，1：检测数据,一般界面，2：重新登录,3:注册页面
    //floorCount:涉及到跳转，需要知道路径层数
    Init: function(flag, floorCount) {
        return init.call(this, flag, floorCount);
    },
    //ajax请求，如果有回调函数，则采用异步的方式，如果没有，则采用非异步的方式返回
    Get: function(className, methodName, data, callback, floorCount) {
        return ajax.call(this, className, methodName, data, 'Get', callback, floorCount, WebMain.WebServerConfig);
    },
    Post: function(className, methodName, data, callback, floorCount) {
        return ajax.call(this, className, methodName, data, 'Post', callback, floorCount, WebMain.WebServerConfig);
    },
    PostPureData: function(className, methodName, data, callback, floorCount) {
        return ajaxPure.call(this, className, methodName, data, 'Post', callback, floorCount, WebMain.WebServerConfig);
    },
    Cookie: function(cookName, cookValue) {
        return cookie.call(this, cookName, cookValue);
    },
    ClearAllCookie: function() {
        return clearAllCookie.call(this);
    },
    SaveLocalData: function(cookName, cookValue) {
        return saveLocalData.call(this, cookName, cookValue);
    },
    GetLocalData: function(cookName) {
        return getLocalData.call(this, cookName);
    },
    //封装sweetalert,
    //title：标题
    //type：类型
    //btnaText：按钮a的文本
    //callbacka：按钮a的回调函数
    //btnbText：按钮b的文本
    //callbackb：按钮b的回调函数
    //btncText：按钮c的文本
    //callbackc：按钮c的回调函数
    Alert: function(title, content, type, btnaText, callbacka, btnbText, callbackb, btncText, callbackc) {
        return alertFunc.call(this, title, content, type, btnaText, callbacka, btnbText, callbackb, btncText, callbackc);
    },
    //获取层级对应的路径
    GetPath: function(floorCount) {
        GetRootPath(floorCount);
    },
}

//初始化,检测数据
function init(flag, floorCount) {
    var result = {}
    checkdata(flag, floorCount);

    //设置默认的提示框
    toastr.options = {
        "closeButton": true,
        "debug": false,
        "progressBar": true,
        "positionClass": "toast-top-right",
        "onclick": null,
        "showDuration": "2000",
        "hideDuration": "1000",
        "timeOut": "5000",
        "extendedTimeOut": "1000",
        "showEasing": "swing",
        "hideEasing": "linear",
        "showMethod": "fadeIn",
        "hideMethod": "fadeOut"
    }

    return result;
}

//检测用户数据
function checkdata(flag, floorCount) {
    var userName = $.cookie("UserName");
    var pwdExpiredTime = $.cookie("PwdExpiredTime");
    var curDate = Date.parse(new Date());
    var rootPath = GetRootPath(floorCount)

    //如果检测数据，那么如果没有用户名，则登录
    if (flag == 1) {
        if (userName == null || userName == "") {
            window.location.href = rootPath + 'login.html';
        } else if (pwdExpiredTime < curDate || pwdExpiredTime == null) {
            //如果有用户名，但是过期了，则重登录
            window.location.href = rootPath + 'lockscreen.html';
        }
    } else if (flag == 0) {
        //如果为登录页面，且密码过期，则重登录
        if (userName != null && userName != "") {
            if (pwdExpiredTime < curDate || pwdExpiredTime == null) {
                //window.location.href = rootPath+'lockscreen.html';
            } else {
                window.location.href = rootPath + 'index.html';
            }
        }
    } else if (flag == 2) {
        if (userName == null || userName == "") {
            window.location.href = rootPath + 'login.html';
        }
    }
}

//ajax请求
function ajax(className, methodName, data, type, callback, floorCount, serverAddress) {
    var result = {}

    var userName = $.cookie("UserName");
    var asyncFlag = !callback ? false : true;
    var rootPath = GetRootPath(floorCount);
    var urlStr = serverAddress + "API/" + className + "/" + methodName;
    var token = $.cookie("Token");

    //调用参数
    var params = {
        UserName: userName,
        Token: token,
        Data: data
    };

    //获取字符串
    var paramStr = JSON.stringify(params);

    var layerIndex = layer.load();
    $.ajax({
        dataType: "text",
        type: type,
        async: asyncFlag,
        url: urlStr,
        data: paramStr,
        success: function(returnInfo) {
            layer.close(layerIndex);

            //如果有回调函数，则调用回调函数来处理数据
            result = returnInfo;
            if (callback) {
                callbackHandle(result, callback, floorCount);
            }
        },
        error: function(request) {
            layer.close(layerIndex);

            toastr.error("提示", "获取数据失败！");
        }
    });

    //如果没有回调函数，则处理数据
    if (!callback)
        return handle(result, floorCount);
}

//ajaxPure请求
function ajaxPure(className, methodName, data, type, callback, floorCount, serverAddress) {
    var result = {}

    var asyncFlag = !callback ? false : true;
    var rootPath = GetRootPath(floorCount);
    var urlStr = serverAddress + "API/" + className + "/" + methodName;

    //调用参数
    var params = {
        Data: data
    };

    //获取字符串
    var paramStr = JSON.stringify(params);

    var layerIndex = layer.load();
    $.ajax({
        dataType: "text",
        type: type,
        async: asyncFlag,
        url: urlStr,
        data: paramStr,
        success: function(returnInfo) {
            layer.close(layerIndex);

            //如果有回调函数，则调用回调函数来处理数据
            var result = JSON.parse(returnInfo);
            if (callback) {
                callback(result)
            }
        },
        error: function(request) {
            layer.close(layerIndex);
        }
    });

    //如果没有回调函数，则处理数据
    if (!callback)
        return result;
}

//处理回调函数的数据
function callbackHandle(returnInfo, callback, floorCount) {
    var data = handle(returnInfo, floorCount);

    if (callback)
        callback(data);
}

//处理返回值
function handle(returnInfo, floorCount) {
    var data = JSON.parse(returnInfo);

    //如果登录超时，直接跳转
    if (data.Status == -8) {
        var rootPath = GetRootPath(floorCount)

        var userName = $.cookie("UserName");
        if (userName == null || userName == "") {
            window.location.href = rootPath + 'login.html';
        } else {
            window.location.href = rootPath + 'lockscreen.html';
        }
        data = {}
    } else {
        //做其他事情
    }

    //如果返回了过期时间
    if (data.AttachData.PwdExpiredTime != null && data.AttachData.PwdExpiredTime != 0) {
        $.cookie("PwdExpiredTime", data.AttachData.PwdExpiredTime, { expires: 30, path: '/' });
    }

    return data;
}

//封装提示框
function alertFunc(title, content, type, btnaText, callbacka, btnbText, callbackb, btncText, callbackc) {
    switch (type) {
        case "success":
            swal({
                title: title,
                text: content,
                type: type
            });
            break;
        case "error":
            swal({
                title: title,
                text: content,
                type: type
            });
            break;
        case "warn":
            swal({
                title: title,
                text: content,
                type: "warning",
                showCancelButton: true,
                cancelButtonText: "取消",
                confirmButtonColor: "#DD6B55",
                confirmButtonText: btnaText,
                closeOnConfirm: false
            }, function() {
                if (callbacka)
                    callbacka();
            });
            break;
        case "customwarn":
            swal({
                title: title,
                text: content,
                type: "warning",
                showCancelButton: true,
                cancelButtonText: btnbText,
                confirmButtonColor: "#DD6B55",
                confirmButtonText: btnaText,
                closeOnConfirm: false,
                closeOnCancel: false
            }, function(isConfirm) {
                if (isConfirm) {
                    if (callbacka)
                        callbacka();
                } else {
                    if (callbackb)
                        callbackb();
                }
            });
            break;
        case "timer":
            swalTimerFunc(title, content, btnaText, callbacka, 3000);
            break;
        default:
            swal({
                title: title,
                text: content
            });
    }
}

//递归函数用来显示时间提示框
function swalTimerFunc(title, content, btnaText, callbacka, i) {
    var currentContent = content.replace("%s", i / 1000 + "s");

    swal({
        title: title,
        text: currentContent,
        timer: 1000,
        type: "success",
        showConfirmButton: true,
        confirmButtonText: btnaText
    }, function(isConfirm) {
        //如果点击了确认，则直接返回
        if (isConfirm) {
            if (callbacka)
                callbacka();
        }

        //继续循环
        i = i - 1000;
        if (i >= 1000) {
            swalTimerFunc(title, content, btnaText, callbacka, i);
        } else {
            if (callbacka)
                callbacka();
        }
    });
}

//存数据
function saveLocalData(cookName, cookValue) {
    localStorage.setItem(cookName, cookValue);
}

//获取数据
function getLocalData(cookName) {
    return localStorage.getItem(cookName);
}

//设置cookie
function cookie(cookName, cookValue) {
    if (typeof(cookValue) != "undefined") {
        if (cookValue == null) {
            $.cookie(cookName, '', { expires: -1, path: '/' });
        } else {
            $.cookie(cookName, cookValue, { expires: 30, path: '/' });
        }
    }
}

//清空cookie
function clearAllCookie() {
    var keys = document.cookie.match(/[^ =;]+(?=\=)/g);
    if (keys) {
        for (var i = keys.length; i--;) {
            document.cookie = keys[i] + '=0;expires=' + new Date(0).toUTCString();
        }
    }
}

//获得根路径
function GetRootPath(floorCount) {
    var rootPath = "";

    if (floorCount == null || floorCount == "undefined" || floorCount == 0) {
        return rootPath;
    } else {
        var tempPath = "../";
        for (var i = 0; i < floorCount; i++) {
            rootPath += tempPath;
        }
    }

    return rootPath;
}

//获取服务器配置
$(function() {
    $.ajax({
        dataType: "text",
        type: "Post",
        async: false,
        url: "/GetConfig",
        success: function(returnInfo) {
            try {
                var returnData = JSON.parse(returnInfo);
                WebMain.FileServerConfig = returnData.Data.FileServerAddress;
                WebMain.WebServerConfig = returnData.Data.GoServerAddress;
                WebMain.ChatServerConfig = returnData.Data.ChatServerAddress;
            } catch (e) {
                toastr.error("提示", "获取服务器失败！");
                return
            }
        },
        error: function(request) {
            toastr.error("提示", "获取服务器失败！");
        }
    });
});