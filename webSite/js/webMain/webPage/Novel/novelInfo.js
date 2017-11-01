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

    //左右键按钮作用
    $(document).keydown(function(e) {
        if (e.keyCode == '37') {
            MovePage(-1)
        } else if (e.keyCode == '39') {
            MovePage(1)
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
        $("#NovelTitle").html(chapterName);
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

            WebMain.SaveLocalData("NovelInfo", returnInfo.Data);

            $("#NovelContent").html("");
            $("#SourceDiv").html("");
            $.each(returnInfo.Data, function(n, value) {
                $("#SourceDiv").append("<button class=\"label-success btn btn-white btn-xs\" type=\"button\" onclick=\"ModifySource()\">" + value.Source + "</button>");

                if (n == returnInfo.Data.length - 1) {
                    $("#NovelContent").html(value.Content);
                    $("#NovelTitle").html(value.Title);
                    $('html,body').animate({ scrollTop: 0 }, 'fast');
                    WebMain.Cookie("ChapterName", value.Title);
                }
            });
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    }, 2);
}

function ModifySource() {
    var sourceText = $(this).Text();
    var novelInfos = WebMain.GetLocalData("NovelInfo");

    $("#SourceDiv button").each(function() {
        $(this).removeClass("label-success");
    });

    $.each(novelInfos, function(n, value) {
        if (value.Source == sourceText) {
            $(this).addClass("label-success");
            $("#NovelContent").html(value.Content);
        }
    });
}