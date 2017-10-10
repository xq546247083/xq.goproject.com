﻿$(document).ready(function () {
    WebMain.Init(0);
});

//回车提交
$(function () {
    $(document).keydown(function (e) {
        if (e.keyCode == "13" ) {
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

    WebMain.Post("SysUser", "Login", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            WebMain.Cookie(
                returnData.Data.UserName,
                returnData.Data.PwdExpiredTime,
                returnData.Data.FullName,
                returnData.Data.Email,
                returnData.Data.Sex,
                returnData.Data.LoginCount,
                returnData.Data.LastLoginTime,
                returnData.Data.LastLoginIP
            );
            window.location.href ='index.html';
        } else {
            toastr.error("提示", returnData.StatusValue);
        }
    });
}