$(document).ready(function() {
    WebMain.Init(1, 2);
    GetPhotoList();
});

//获取照片列表
function GetPhotoList() {
    var userName = $.cookie("UserName");
    var token = $.cookie("Token");
    var urlStr = WebMain.FileServerConfig + "API/Photo/GetUserPhotos";

    //方法参数
    var data = new Array();
    data[0] = userName;

    //调用参数
    var params = {
        UserName: userName,
        Token: token,
        Data: data
    };

    //获取字符串
    var paramStr = JSON.stringify(params);

    var layerIndex = layer.load();
    $.ajax({
        dataType: "text",
        type: "Post",
        async: false,
        url: urlStr,
        data: paramStr,
        success: function(returnInfo) {
            layer.close(layerIndex);
            handle(returnInfo);
        },
        error: function(request) {
            layer.close(layerIndex);
        }
    });
}

function handle(returnInfo) {
    $("#imgList").html("");
    var returnData = JSON.parse(returnInfo);
    //遍历返回的元素
    $.each(returnData.Data, function(n, value) {
        var photoName = "照片" + (n + 1);
        var imgUrl = (WebMain.FileServerConfig + value.DirName + value.FileName).replace("./", "");
        var modTimeStr = value.ModName.substr(0, 16).replace("T", " ");
        var imgStr = "<a href=\"" + imgUrl + "\" title=\"" + photoName + "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;" + modTimeStr + "\" data-gallery=\"\"><img  src=\"" + imgUrl + "\"></a>"
        $("#imgList").append(imgStr);
    });
}