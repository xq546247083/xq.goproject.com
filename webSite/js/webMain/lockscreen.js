$(document).ready(function() {
    WebMain.Init(2);

    //进入这个页面，直接判定为过期
    var userName = $.cookie("UserName");
    $("#UserName").html(userName);
    WebMain.Cookie(userName, null);

    //回车提交
    $(document).keydown(function(e) {
        if (e.keyCode == '13') {
            ReLogin();
        }
    })
});

//获取接口文档
function ReLogin() {
    var userName = $.cookie("UserName");
    if (userName == null || userName == "") {
        window.location.href = 'login.html';
    }

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
            window.location.href = 'index.html';
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    });
}