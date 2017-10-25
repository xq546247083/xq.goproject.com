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

//上传代码
$(document).ready(function() {
    var uploadFlag = false;
    var uploader = WebUploader.create({
        pick: {
            id: "#filePicker",
            label: "点击上传头像",
            multiple: false,
        },
        auto: true,
        accept: {
            title: "Images",
            extensions: "gif,jpg,jpeg,bmp,png",
            mimeTypes: "image/jpg,image/jpeg,image/png"
        },
        // compress: false,
        compress: {
            width: 600,
            height: 400,

            // 图片质量，只有type为`image/jpeg`的时候才有效。
            quality: 90,

            // 是否允许放大，如果想要生成小图的时候不失真，此选项应该设置为false.
            allowMagnify: false,

            // 是否允许裁剪。
            crop: false,

            // 是否保留头部meta信息。
            preserveHeaders: true,

            // 如果发现压缩后文件大小比原来还大，则使用原来图片
            // 此属性可能会影响图片自动纠正功能
            noCompressIfLarger: false,

            // 单位字节，如果图片大小小于此值，不会采用压缩。
            compressSize: 0
        },
        fileNumLimit: 1,
        fileSizeLimit: 5242880,
        fileSingleSizeLimit: 5242880,
        chunked: true, //开启分片上传  
        chunkSize: 1024 * 100, // 如果要分片，分多大一片？默认大小为5M  
        chunkRetry: 3, //如果某个分片由于网络问题出错，允许自动重传多少次  
        threads: 1, //上传并发数。允许同时最大上传进程数[默认值：3]  
        duplicate: false, //是否重复上传（同时选择多个一样的文件），true可以重复上传  
        prepareNextFile: true, //上传当前分片时预处理下一分片  
        //-------------------------设置上传的服务器地址----------------------
        server: WebMain.FileServerConfig + "APIFromFile/UploadPhoto",
        formData: {
            UserName: "",
            Token: "",
            UploadTime: "",
            PicName: "",
        },
    });

    uploader.on('uploadBeforeSend', function(obj, data, headers) {
        //赋值
        data.UploadTime = new Date().getTime();
        data.PicName = "Album";
        data.UserName = $.cookie("UserName");
        data.Token = $.cookie("Token");
        $("#filePicker").hide();
    });

    uploader.on('uploadError', function(e) {
        toastr.error("提示", "上传失败");
    });

    uploader.on('uploadComplete', function(file, returnInfo) {
        toastr.success("提示", "上传完成");
    });

    uploader.on('uploadAccept', function(file, returnInfo) {
        //如果没修改过用户的头像，则更新头像
        if (!uploadFlag) {
            var photoUrl = returnInfo.Data.DirName + returnInfo.Data.FileName;
            if (photoUrl != null && photoUrl != "") {
                //方法参数
                var data = new Array();
                data[0] = $.cookie("UserName");;
                data[1] = photoUrl;

                WebMain.Post("SysUser", "UpdatePhoto", data, function(returnInfo) {
                    if (returnInfo == {}) return;
                    if (returnInfo.Status == 0) {
                        $("#HeadImg").attr("src", WebMain.FileServerConfig + returnInfo.Data.HeadImgage);
                        WebMain.Cookie("HeadImgage", returnInfo.Data.HeadImgage);
                    } else {
                        toastr.error("提示", returnInfo.StatusValue);
                    }
                });

            }

            uploadFlag = true;
        }
    });
});