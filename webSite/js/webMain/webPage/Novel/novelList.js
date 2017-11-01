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

    //小说列表点击
    $(document).on("click", "li.Novel", function(e) {
        WebMain.Cookie("NovelName", $(this).text().substring(2));
        window.location.href = 'chapterList.html';
    });

    //左右键按钮作用
    $(document).keydown(function(e) {
        if (e.keyCode == '13') {

        }
    });

    GetNovelList();
});

//获取小说列表
function GetNovelList() {
    //方法参数
    var data = new Array();

    var classList = new Array("warning-element", "success-element", "info-element", "danger-element", "warning-element");

    WebMain.PostPureData("Novel", "GetNovelList", data, function(returnInfo) {
        if (returnInfo == {}) return;

        if (returnInfo.Status == 0) {
            //遍历返回的元素
            $.each(returnInfo.Data, function(n, value) {
                $("#content").append("<li class =\"Novel " + classList[n % 5] + "\" >&nbsp;&nbsp;" + value + "</li>");
            });
        } else {
            toastr.error("提示", returnInfo.StatusValue);
        }
    }, 2);
}