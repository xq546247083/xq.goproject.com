$(document).ready(function() {
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
    WebMain.ClearAllCookie();

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