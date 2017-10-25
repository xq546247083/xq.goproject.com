$(document).ready(function() {
    WebMain.Init(1, 2);

    var fullName = $.cookie("FullName");
    var email = $.cookie("Email");
    var sex = $.cookie("Sex");
    var headImgage = $.cookie("HeadImgage");
    var loginCount = $.cookie("LoginCount");
    var lastLoginTime = $.cookie("LastLoginTime");

    //设置页面信息
    $("#hName").html(fullName);
    if (headImgage != null && headImgage != "") {
        $("#HeadImg").attr("src", WebMain.FileServerConfig + headImgage);
    }

    var curLastLoginTime = lastLoginTime.substr(0, 16).replace("T", " ");
    $("#info").prepend("<p>最后登录时间:" + curLastLoginTime + "</p>");
    $("#info").prepend("<p>登录次数:" + loginCount + "</p>");
    $("#info").prepend("<p>邮箱：" + email + "</p>");
    $("#info").prepend("<p>性别:" + (sex ? "男" : "女") + "</p>");
    $("#info").prepend("<h2 class=\"media-heading\">" + fullName + "</h2><br/>");


});