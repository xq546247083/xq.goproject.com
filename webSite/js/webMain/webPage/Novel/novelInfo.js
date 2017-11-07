$(document).ready(function() {
    //设置默认的提示框
    toastr.options = {
        "closeButton": true,
        "debug": false,
        "progressBar": true,
        "positionClass": "toast-top-right",
        "onclick": null,
        "showDuration": "2000",
        "hideDuration": "1000",
        "timeOut": "5000",
        "extendedTimeOut": "1000",
        "showEasing": "swing",
        "hideEasing": "linear",
        "showMethod": "fadeIn",
        "hideMethod": "fadeOut"
    }

    //修改首页
    $(".gohome").html('<a class="animated bounceInUp" href="novelList.html" title="返回首页"><i class="fa fa-home"></i></a>');
    $("#NovelContent").css("background-color", "");

    $("#NovelContent").css("color", $.cookie("FontColor"));

    //左右键按钮作用
    $(document).keydown(function(e) {
        if (e.keyCode == '37') {
            MovePage(-1)
        } else if (e.keyCode == '39') {
            MovePage(1)
        } else if (e.keyCode == '13') {
            window.location.href = 'chapterList.html';
        }
    });

    MovePage(0)
});

function MovePage(flag) {
    //设置页面属性
    var novelName = $.cookie("NovelName");
    if (novelName == null || novelName == "") {
        window.location.href = 'novelList.html';
    }

    var chapterName = $.cookie("ChapterName");
    if (chapterName == null || chapterName == "") {
        window.location.href = 'chapterList.html';
    } else {
        $("title").html(chapterName);
    }

    GetNovelInfo(novelName, chapterName, flag);
}
//获取小说
function GetNovelInfo(novelName, chapterName, flag) {
    //方法参数
    var data = new Array();
    data[0] = novelName;
    data[1] = chapterName;
    data[2] = flag;

    var classList = new Array("warning-element", "success-element", "info-element", "danger-element", "warning-element");

    WebMain.PostPureData("Novel", "GetNovelInfo", data, function(returnInfo) {
        if (returnInfo == {}) return;

        if (returnInfo.Status == 0) {
            if (returnInfo.Data == null || returnInfo.Data.length == 0) {
                return
            }

            WebMain.SaveLocalData("NovelInfo", JSON.stringify(returnInfo.Data));

            $("#NovelContent").html("");
            $("#SourceDiv").html("");
            $.each(returnInfo.Data, function(n, value) {
                if (n == returnInfo.Data.length - 1) {
                    $("#SourceDiv").append("<button class=\"label-success btn btn-white btn-xs\" type=\"button\" onclick=\"ModifySource(this)\">" + value.Source + "</button>&nbsp;&nbsp;");
                    $("#NovelContent").html(value.Content);
                    $("title").html(value.Title);
                    $('html,body').animate({ scrollTop: 0 }, 'fast');
                    WebMain.Cookie("ChapterName", value.Title);
                } else {
                    $("#SourceDiv").append("<button class=\" btn btn-white btn-xs\" type=\"button\" onclick=\"ModifySource(this)\">" + value.Source + "</button>&nbsp;&nbsp;");
                }
            });
            var fontSize = $.cookie("FontSize");
            $("#NovelContent p").css("font-size", fontSize);
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    }, 2);
}

//修改源
function ModifySource(e) {
    var sourceText = $(e).text();
    var novelInfos = JSON.parse(WebMain.GetLocalData("NovelInfo"));

    $("#SourceDiv button").each(function() {
        $(this).removeClass("label-success");
    });

    $.each(novelInfos, function(n, value) {
        if (value.Source == sourceText) {
            $(e).addClass("label-success");
            $("#NovelContent").html(value.Content);
            return
        }
    });

    var fontSize = $.cookie("FontSize");
    $("#NovelContent p").css("font-size", fontSize);
}

//修改字体大小
function ChangeFont(flag) {
    var originalFontSizeStr = $("#NovelContent p").css("font-size").replace(/[^0-9]/ig, "");;

    var originalFontSize = parseInt(originalFontSizeStr);
    if (flag) {
        originalFontSize += 2;
    } else {
        originalFontSize -= 2;
    }

    $("#NovelContent p").css("font-size", originalFontSize + "px");
    WebMain.Cookie("FontSize", originalFontSize + "px");
}