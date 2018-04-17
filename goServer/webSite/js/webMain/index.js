$(document).ready(function() {
    WebMain.Init(1);

    var fullName = $.cookie("FullName");
    var email = $.cookie("Email");
    var sex = $.cookie("Sex");
    var headImage = $.cookie("HeadImage");

    //设置页面信息
    $("#userNameSpan").html(fullName);
    $(document).attr("title", fullName + "的主页");
    if (headImage != null && headImage != "") {
        $("#HeadImg").attr("src", WebMain.FileServerConfig + headImage);
    }

    GetMenu();
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

//锁屏
function LockScreen() {
    var userName = $.cookie("UserName");
    var headImage = $.cookie("HeadImage");
    if (userName == null || userName == "") {
        window.location.href = 'lockscreen.html';
    }

    //方法参数
    var data = new Array();
    data[0] = userName;

    //先退出聊天服务器
    ChatMain.SendMessage("Logout", "", "");
    WebMain.Post("SysUser", "LoginOut", data, function(returnInfo) {
        WebMain.ClearAllCookie();
        WebMain.Cookie("UserName", userName);
        WebMain.Cookie("HeadImage", headImage);
        window.location.href = 'lockscreen.html';
    });
}


//退出
function LoginOut() {
    var userName = $.cookie("UserName");
    if (userName == null || userName == "") {
        window.location.href = 'login.html';
    }

    //方法参数
    var data = new Array();
    data[0] = userName;

    //先退出聊天服务器
    ChatMain.SendMessage("Logout", "", "");
    WebMain.Post("SysUser", "LoginOut", data, function(returnInfo) {
        WebMain.ClearAllCookie();
        window.location.href = 'login.html';
    });
}