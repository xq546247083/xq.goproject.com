$(document).ready(function () {
    WebMain.Init(2);

    //进入这个页面，直接判定为过期
    var userName = $.cookie("UserName");
    $("#UserName").html(userName);
    WebMain.Cookie(userName, null);

    //回车提交
    $(document).keydown(function (e) {
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

    WebMain.Post("SysUser", "Login", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            WebMain.Cookie(
                returnData.Value.UserName,
                returnData.Value.PwdExpiredTime,
                returnData.Value.FullName,
                returnData.Value.Email,
                returnData.Value.Sex,
                returnData.Value.LoginCount,
                returnData.Value.LastLoginTime,
                returnData.Value.LastLoginIP
            );
            window.location.href = 'index.html';
        } else {
            toastr.error("提示", returnData.StatusValue);
        }
    });
}