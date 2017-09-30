$(document).ready(function () {
    WebMain.Init(3);

    $(".i-checks").iCheck({
        checkboxClass: "icheckbox_square-green",
        radioClass: "iradio_square-green"
    });
});

//验证
$.validator.setDefaults({
    highlight: function (e) {
        $(e).closest(".form-group").removeClass("has-success").addClass("has-error")
    },
    success: function (e) {
        e.closest(".form-group").removeClass("has-error").addClass("has-success")
    },
    errorElement: "span", errorPlacement: function (e, r) {
        e.appendTo(r.is(":radio") || r.is(":checkbox") ? r.parent().parent().parent() : r.parent())
    },
    errorClass: "help-block m-b-none", validClass: "help-block m-b-none"
}), $().ready(function () {
    var e = "<i class='fa fa-times-circle'></i> ";
    $("#registerForm").validate({
        rules: {
            userName: {
                required: !0, minlength: 4
            },
            userPassword1: {
                required: !0, minlength: 8
            },
            userPassword2: {
                required: !0, minlength: 8, equalTo: "#userPassword1"
            },
            email: {
                required: !0, email: true
            },
            identifyCode: {
                rangelength: [6,6]
            },
        },
        messages:
        {
            userName: { required: e + "请输入您的用户名", minlength: e + "用户名必须四个字符以上" },
            userPassword1: { required: e + "请输入您的密码", minlength: e + "密码必须八个字符以上" },
            userPassword2: { required: e + "请再次输入密码", minlength: e + "密码必须八个字符以上", equalTo: e + "两次输入的密码不一致" },
            email: { required: e + "请输入您的E-mail", email: e +"请输入正确的邮箱" },
            identifyCode: { rangelength: e + "验证码长度为六位" }
        }
    })
});


//回车提交
$(function () {
    $(document).keydown(function (e) {
        if (e.keyCode == "13") {
            Register()
        }
    })
})

function Identify() {
    var email = $("#email").val();
    var data = new Array();
    data[0] = email;
    data[1] = 0;

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

//注册
function Register() {
    var userPassword1 = $("#userPassword1").val();
    var userPassword2 = $("#userPassword2").val();
    if (userPassword1 != userPassword2) {
        toastr.error("提示", "密码不相同");
        return;
    }
    var protocolCb = $("#protocolCb").is(':checked');
    if (protocolCb == false) {
        toastr.warning("提示", "协议未勾选");
        return;
    }

    var userName = $("#userName").val();
    var fullName = $("#userName").val();
    var email = $("#email").val();
    var identifyCode = $("#identifyCode").val();

    var pwd = md5(userPassword1);

    //方法参数
    var data = new Array();
    data[0] = userName;
    data[1] = pwd;
    data[2] = fullName;
    data[3] = 1;
    data[4] = email;
    data[5] = identifyCode;

    WebMain.Post("SysUser", "Register", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            WebMain.Alert("注册成功", "点击OK跳转登录页面，%s后自动跳转登录页面..", "timer", "OK", function () {
                window.location.href = 'login.html';
            });
        } else {
            toastr.error("提示", returnData.StatusValue);
        }
    });
}