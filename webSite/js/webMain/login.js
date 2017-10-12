﻿$(document).ready(function() {
    WebMain.Init(0);
});

//回车提交
$(function() {
    $(document).keydown(function(e) {
        if (e.keyCode == "13") {
            Login()
        }
    })
})

//获取接口文档
function Login() {
    //清空cookie
    WebMain.Cookie(null, null);

    var userName = $("#userName").val();
    var userPassword = $("#userPassword").val();
    var pwd = md5(userPassword);

    //方法参数
    var data = new Array();
    data[0] = userName;
    data[1] = pwd;

    WebMain.Post("SysUser", "Login", data, function(returnInfo) {
        if (returnInfo == {}) return;

        if (returnInfo.Status == 0) {
            WebMain.Cookie(
                returnInfo.Data.UserName,
                returnInfo.Data.PwdExpiredTime,
                returnInfo.Data.FullName,
                returnInfo.Data.Email,
                returnInfo.Data.Sex,
                returnInfo.Data.LoginCount,
                returnInfo.Data.LastLoginTime,
                returnInfo.Data.LastLoginIP
            );
            WebMain.CookieOneKey("Token", returnInfo.Data.Token);
            window.location.href = 'index.html';
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    });
}