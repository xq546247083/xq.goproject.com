package emailTool

import (
	"testing"

	"xq.goproject.com/commonTools/configTool"
)

func TestSendMail(t *testing.T) {
	SetSenderInfo(configTool.EmailHost, 465, configTool.EmailName, configTool.EmailAddress, configTool.EmailPass)
	err := SendMail([]string{"295787943@qq.com"}, "邮件发送测试", "<h1>这是邮件正文</h1>", true, nil)
	if err != nil {
		t.Error(err)
	}
}
