$(document).ready(function () {
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
function GetMenu(){
    var userName = $.cookie("UserName");

    //方法参数
    var data = new Array();
    data[0] = userName;

    WebMain.Post("SysMenu", "GetInfo", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            $("#side-menu").append(returnData.Value.MenuScript);
        } else {
            toastr.error("提示", returnData.StatusValue);
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

    WebMain.Post("SysUser", "LoginOut", data, function (returnData) {
        WebMain.Cookie(null, null, null, null, null, null, null, null
        );
        window.location.href = 'login.html';
    });
}

//获取接口文档
function LockScreen() {
    var userName = $.cookie("UserName");
    WebMain.Cookie(userName, null);
    window.location.href = 'lockscreen.html';
}