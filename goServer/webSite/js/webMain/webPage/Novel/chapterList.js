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

    //设置页面属性
    var novelName = $.cookie("NovelName");
    if (novelName == null || novelName == "") {
        window.location.href = 'novelList.html';
    } else {
        $("#chapterTitle").html(novelName);
        $("title").html(novelName);
    }

    //小说列表点击
    $(document).on("click", "li.Chapter", function(e) {
        WebMain.Cookie("ChapterName", $(this).text().substring(2));
        window.location.href = 'novelInfo.html';
    });

    //左右键按钮作用
    $(document).keydown(function(e) {
        if (e.keyCode == '37') {
            MovePage(false)
        } else if (e.keyCode == '39') {
            MovePage(true)
        } else if (e.keyCode == '13') {
            window.location.href = 'novelList.html';
        }
    });

    var pageNum = $.cookie(novelName + "NovelPageNum");
    if (pageNum == null) {
        pageNum = 1;
    }

    $("#pageNum").html(pageNum);
    GetChapterList(novelName, pageNum);
});

// 移动页面
function MovePage(flag) {
    var novelName = $.cookie("NovelName");

    var pageNum = $.cookie(novelName + "NovelPageNum");
    if (pageNum == null) {
        pageNum = 1;
    }

    if (flag) {
        pageNum++
    } else {
        pageNum--
        if (pageNum <= 0) {
            pageNum = 1
        }
    }



    GetChapterList(novelName, pageNum);
}

//获取小说列表
function GetChapterList(novelName, pageNum) {
    //方法参数
    var data = new Array();
    data[0] = novelName;
    data[1] = pageNum;

    var classList = new Array("warning-element", "success-element", "info-element", "danger-element", "warning-element");

    WebMain.PostPureData("Novel", "GetChapterList", data, function(returnInfo) {
        if (returnInfo == {}) { return };

        if (returnInfo.Status == 0) {
            if (returnInfo.Data == null || returnInfo.Data.length == 0) {
                return
            }

            //遍历返回的元素
            WebMain.Cookie(novelName + "NovelPageNum", pageNum);
            $("#pageNum").html(pageNum);
            $("#content").html("");
            $.each(returnInfo.Data, function(n, value) {
                $("#content").append("<li class =\"Chapter " + classList[n % 5] + "\" >&nbsp;&nbsp;" + value + "</li>");
            });
            $('html,body').animate({ scrollTop: 0 }, 'fast');
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    }, 2);
}