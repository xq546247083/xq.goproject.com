$(document).ready(function() {
    WebMain.Init(1);

    var fullName = $.cookie("FullName");
    var email = $.cookie("Email");
    var sex = $.cookie("Sex");

    //设置页面信息
    $("#userNameSpan").html(fullName);
    $(document).attr("title", fullName + "的主页");

    GetMenu()
});

//获取菜单信息
function GetMenu() {
    var userName = $.cookie("UserName");

    //方法参数
    var data = new Array();
    data[0] = userName;

    WebMain.Post("SysMenu", "GetInfo", data, function(returnInfo) {
        if (returnInfo == {}) return;

        if (returnInfo.Status == 0) {
            $("#side-menu").append(returnInfo.Data.MenuScript);
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    });
}

//获取接口文档
function LoginOut() {
    var userName = $.cookie("UserName");
    if (userName == null || userName == "") {
        window.location.href = 'login.html';
    }

    //方法参数
    var data = new Array();
    data[0] = userName;

    WebMain.Post("SysUser", "LoginOut", data, function(returnInfo) {
        WebMain.Cookie(null, null, null, null, null, null, null, null);
        window.location.href = 'login.html';
    });
}

//获取接口文档
function LockScreen() {
    var userName = $.cookie("UserName");
    WebMain.Cookie(userName, null);
    window.location.href = 'lockscreen.html';
}