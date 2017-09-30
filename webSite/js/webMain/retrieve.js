$(document).ready(function () {
    WebMain.Init(3);
});

//回车提交
$(function () {
    $(document).keydown(function (e) {
        if (e.keyCode == "13") {
            Retrieve()
        }
    })
})


function Identify() {
    var email = $("#email").val();
    var data = new Array();
    data[0] = email;
    data[1] = 1;

    WebMain.Post("SysUser", "Identify", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            toastr.success("提示", "邮件已发送");

            //一分钟可点击一次发送邮件
            $("#identifyBtn").attr('disabled', true);
            setTimeout(function () {
                $("#identifyBtn").attr('disabled', false);
            }, 60000);
        } else {
            toastr.error("提示", returnData.StatusValue);
        }
    });
}

//找回密码
function Retrieve() {
    var userPassword1 = $("#userPassword1").val();
    var userPassword2 = $("#userPassword2").val();
    if (userPassword1 != userPassword2) {
        toastr.error("提示", "密码不相同");
        return;
    }

    var email = $("#email").val();
    var identifyCode = $("#identifyCode").val();

    var pwd = md5(userPassword1);

    //方法参数
    var data = new Array();
    data[0] = pwd;
    data[1] = email;
    data[2] = identifyCode;

    WebMain.Post("SysUser", "Retrieve", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            WebMain.Alert("成功找回密码", "点击OK跳转登录页面，%s后自动跳转登录页面..", "timer", "OK", function () {
                window.location.href = 'login.html';
            });
        } else {
            toastr.error("提示", returnData.StatusValue);
        }
    });
}