$(document).ready(function () {
    WebMain.Init(1,2);
    GetBlogList(2, 1)
});

//获取博客列表
function GetBlogList(blogType, status) {
    var userName = $.cookie("UserName");

    //方法参数
    var data = new Array();
    data[0] = userName;
    data[1] = blogType;
    data[2] = status;
    data[3] = "";

    WebMain.Post("UBlog", "GetBlogList", data, function (returnData) {
        if (returnData == {}) return;

        if (returnData.Status == 0) {
            var e1 = "<tr class='read'><td class='check-mail'><input type='checkbox' class='i-checks'></td><td class='mail- ontact'><a href='#'>";
            var e2 = "</a></td><td class='mail-subject'><a href='BlogView.html'>";
            var e3 = "</a></td><td class=''><i class='fa fa- paperclip'></i></td><td class='text- right mail- date'>";
            var e4 = "</td></tr>";

            //构造元素
            var htmlContent = "";
            $.each(returnData.Value, function (index, item) {
                htmlContent += e1 + item.Title + e2 + item.Content + e3 + item.ReDate + e4;
            });

            $("#bloglistTD").append(htmlContent);

            //构造复选框样式
            $(".i-checks").iCheck({
                checkboxClass: "icheckbox_square-green", radioClass: "iradio_square-green",
            })
        } else {
            toastr.error("提示", returnData.StatusValue);
        }
    },2);
}


$(document).on("click", "#blogQueryBtn", function (e) {
    var name = $(e).attr("name");
    if (name == "note") {
        GetBlogList(1, 1);
    } else if (name == "blog") {
        GetBlogList(2, 1);
    } else if (name == "draft") {
        GetBlogList(-1, 0);
    } else if (name == "dustbin") {
        GetBlogList(-1, 2);
    }
});