$(document).ready(function() {
    WebMain.Init(2);

    var userName = $.cookie("UserName");
    var headImage = $.cookie("HeadImage");

    $("#UserName").html(userName);
    if (headImage != null && headImage != "") {
        $("#HeadImg").attr("src", WebMain.FileServerConfig + headImage);
    }

    //进入这个页面，直接判定为过期
    WebMain.ClearAllCookie();
    WebMain.Cookie("UserName", userName);
    WebMain.Cookie("HeadImage", headImage);

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
            WebMain.Cookie("UserName", returnInfo.Data.UserName);
            WebMain.Cookie("PwdExpiredTime", returnInfo.Data.PwdExpiredTime);
            WebMain.Cookie("FullName", returnInfo.Data.FullName);
            WebMain.Cookie("Email", returnInfo.Data.Email);
            WebMain.Cookie("Sex", returnInfo.Data.Sex);
            WebMain.Cookie("LoginCount", returnInfo.Data.LoginCount);
            WebMain.Cookie("LastLoginTime", returnInfo.Data.LastLoginTime);
            WebMain.Cookie("LastLoginIP", returnInfo.Data.LastLoginIP);
            WebMain.Cookie("HeadImage", returnInfo.Data.HeadImage);
            WebMain.Cookie("Token", returnInfo.Data.Token);
            window.location.href = 'index.html';
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    });
}