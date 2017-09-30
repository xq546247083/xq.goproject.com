var WebMain = {
    //----------------------------------------------一些配置----------------------------------------------
    //业务服务器配置
    WebServerConfig:"http://xiaohe.nat123.cc:48014/",

    //文件服务器配置
    FileServerConfig:"http://107.151.172.51:8882/",

    //----------------------------------------------一些通用方法----------------------------------------------
    //初始化,检测数据
    //flag: 0：登录页面，1：检测数据,一般界面，2：重新登录,3:注册页面
    //floorCount:涉及到跳转，需要知道路径层数
    Init: function (flag,floorCount) {
        return init.call(this, flag,floorCount);
    },
    //ajax请求，如果有回调函数，则采用异步的方式，如果没有，则采用非异步的方式返回
    Get: function (className, methodName, data, callback,floorCount) {
        return ajax.call(this, className, methodName, data, 'Get', callback,floorCount);
    },
    Post: function (className, methodName, data, callback,floorCount) {
        return ajax.call(this, className, methodName, data, 'Post', callback,floorCount);
    },
    Cookie: function (userName, pwdExpiredTime, fullName, email, sex, loginCount, lastLoginTime, lastLoginIP) {
        return cookie.call(this, userName, pwdExpiredTime, fullName, email, sex, loginCount, lastLoginTime, lastLoginIP);
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
    Alert: function (title, content, type, btnaText, callbacka, btnbText, callbackb, btncText, callbackc) {
        return alertFunc.call(this, title, content, type, btnaText, callbacka, btnbText, callbackb, btncText, callbackc);
    },
    //获取层级对应的路径
    GetPath: function(floorCount){
        GetRootPath(floorCount);
    },
}

//初始化,检测数据
function init(flag,floorCount) {
    var result = {}
    checkdata(flag,floorCount);

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
            window.location.href =rootPath+'login.html';
        } else if (pwdExpiredTime < curDate || pwdExpiredTime == null) {
            //如果有用户名，但是过期了，则重登录
            window.location.href =rootPath+'lockscreen.html';
        }
    } else if (flag == 0) {
        //如果为登录页面，且密码过期，则重登录
        if (userName != null && userName != "") {
            if (pwdExpiredTime < curDate || pwdExpiredTime == null) {
                //window.location.href = rootPath+'lockscreen.html';
            } else {
                window.location.href = rootPath+'index.html';
            }
        }
    } else if (flag == 2) {
        if (userName == null || userName == "") {
            window.location.href = rootPath+'login.html';
        }
    }
}

//ajax请求
function ajax(className, methodName, data, type, callback,floorCount) {
    var result = {}

    var userName = $.cookie("UserName");
    var asyncFlag = !callback ? false : true;
    var rootPath = GetRootPath(floorCount)

    //调用参数
    var params = {
        ClassName: className,
        MethodName: methodName,
        UserName: userName,
        Data: data
    };

    var paramStr = JSON.stringify(params);
    var layerIndex=layer.load();
    $.ajax({
        dataType: "text",
        type: type,
        async: asyncFlag,
        url: WebMain.WebServerConfig+"API/ClientHandler.ashx",
        data: paramStr,
        success: function (returnData) {
            layer.close(layerIndex);
            
            //如果有回调函数，则调用回调函数来处理数据
            result = returnData;
            if (callback) {
                callbackHandle(result, callback,floorCount);
            }
        },
        error: function (request) {
            layer.close(layerIndex);

            if (request.status == 500) {
                window.location.href = rootPath+'500.html';
            } else {
                window.location.href = rootPath+'404.html';
            }
        }
    });

    //如果没有回调函数，则处理数据
    if (!callback)
        return handle(result,floorCount);
}

//处理回调函数的数据
function callbackHandle(returnData, callback,floorCount) {
    var data = handle(returnData,floorCount);

    if (callback)
        callback(data);
}

//处理返回值
function handle(returnData,floorCount) {
    var data = JSON.parse(returnData);

    //如果登录超时，直接跳转
    if (data.Status == 7) {
        var rootPath = GetRootPath(floorCount)

        var userName = $.cookie("UserName");
        if (userName == null || userName == "") {
            window.location.href = rootPath+'login.html';
        } else {
            window.location.href = rootPath+'lockscreen.html';
        }
        data = {}
    } else {
        //做其他事情
    }

    //如果返回了过期时间
    if (data.PwdExpiredTime != null && data.PwdExpiredTime != 0) {
        $.cookie("PwdExpiredTime", data.PwdExpiredTime, { expires: 30, path: '/' });
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
            }, function () {
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
            }, function (isConfirm) {
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
function swalTimerFunc(title, content,  btnaText, callbacka, i) {
    var currentContent = content.replace("%s", i / 1000 + "s");

    swal({
        title: title,
        text: currentContent,
        timer: 1000,
        type: "success",
        showConfirmButton: true,
        confirmButtonText: btnaText
    }, function (isConfirm) {
        //如果点击了确认，则直接返回
        if (isConfirm) {
            if (callbacka)
                callbacka();
        }

        //继续循环
        i = i - 1000;
        if (i >= 1000) {
            swalTimerFunc(title, content,  btnaText, callbacka, i);
        } else {
            if (callbacka)
                callbacka();
        }
    });
}

//设置cookie
function cookie(userName, pwdExpiredTime, fullName, email, sex, loginCount, lastLoginTime, lastLoginIP) {
    if (typeof (userName) != "undefined") {
        if (userName == null) {
            $.cookie('UserName', '', { expires: -1, path: '/' });
        } else {
            $.cookie("UserName", userName, { expires: 30, path: '/' });
        }
    }

    if (typeof (pwdExpiredTime) != "undefined") {
        if (pwdExpiredTime == null) {
            $.cookie('PwdExpiredTime', '', { expires: -1, path: '/' });
        } else {
            $.cookie("PwdExpiredTime", pwdExpiredTime, { expires: 30, path: '/' });
        }
    }

    if (typeof (fullName) != "undefined") {
        if (fullName == null) {
            $.cookie('FullName', '', { expires: -1, path: '/' });
        } else {
            $.cookie("FullName", fullName, { expires: 30, path: '/' });
        }
    }

    if (typeof (email) != "undefined") {
        if (email == null) {
            $.cookie('Email', '', { expires: -1, path: '/' });
        } else {
            $.cookie("Email", email, { expires: 30, path: '/' });
        }
    }

    if (typeof (sex) != "undefined") {
        if (sex == null) {
            $.cookie('Sex', '', { expires: -1, path: '/' });
        } else {
            $.cookie("Sex", sex, { expires: 30, path: '/' });
        }
    }

    if (typeof (loginCount) != "undefined") {
        if (loginCount == null) {
            $.cookie('LoginCount', '', { expires: -1, path: '/' });
        } else {
            $.cookie("LoginCount", loginCount, { expires: 30, path: '/' });
        }
    }

    if (typeof (lastLoginTime) != "undefined") {
        if (lastLoginTime == null) {
            $.cookie('LastLoginTime', '', { expires: -1, path: '/' });
        } else {
            $.cookie("LastLoginTime", lastLoginTime, { expires: 30, path: '/' });
        }
    }
    if (typeof (lastLoginIP) != "undefined") {
        if (pwdExpiredTime == null) {
            $.cookie('LastLoginIP', '', { expires: -1, path: '/' });
        } else {
            $.cookie("LastLoginIP", lastLoginIP, { expires: 30, path: '/' });
        }
    }
}

//获得根路径
function GetRootPath(floorCount){
    var rootPath="";

    if(floorCount==null||floorCount=="undefined"||floorCount==0){
        return rootPath;
    }else {
        var tempPath="../";
        for (var i=0;i<floorCount;i++)
        {
            rootPath+=tempPath;
        }
    }

    return rootPath;
}