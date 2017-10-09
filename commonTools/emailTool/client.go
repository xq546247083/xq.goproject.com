package emailTool

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"path/filepath"
	"strings"
)

// SMTPClient实现
type simpleClient struct {
	mailHost   string
	mailPort   int32
	senderName string
	senderAddr string
	senderPwd  string
}

var (
	// 邮箱客户端
	emailClient = new(simpleClient)
)

// SetSenderInfo 设置发送者消息
func SetSenderInfo(mailHost string, mailPort int32, name string, address string, password string) {
	emailClient.mailHost = mailHost
	emailClient.mailPort = mailPort
	emailClient.senderName = name
	emailClient.senderAddr = address
	emailClient.senderPwd = password
}

// SendMail 发送邮件:
//  mailTo 接收方列表
//  subject 主题
//  body 正文
//  isHTMLBody 正文是否html格式
//  attachFiles 附件
func SendMail(mailTo []string, subject, body string, isHTMLBody bool, attachFiles []string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         emailClient.mailHost,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", emailClient.mailHost, emailClient.mailPort), tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	// 创建smtp.Client
	c, err := smtp.NewClient(conn, emailClient.mailHost)
	if err != nil {
		return err
	}

	// 验证信息
	auth := smtp.PlainAuth("", emailClient.senderAddr, emailClient.senderPwd, emailClient.mailHost)
	if err = c.Auth(auth); err != nil {
		return err
	}

	// 发送方
	from := mail.Address{emailClient.senderName, emailClient.senderAddr}
	// 接收方
	to := make([]string, 0, len(mailTo))
	for _, v := range mailTo {
		to = append(to, "<"+v+">")
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		return err
	}

	for _, v := range mailTo {
		if err = c.Rcpt(v); err != nil {
			return err
		}
	}

	// 边界
	boundary := "a40acf3c8b7200fc6b04c2f1b3da"

	buff := bytes.NewBuffer(nil)

	// 写入基本信息
	buff.WriteString(fmt.Sprintf("From: %s\r\n", from.String()))
	buff.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(to, ", ")))
	buff.WriteString(fmt.Sprintf("Subject: %s\r\n", subject))

	// 写入邮件头部信息
	if len(attachFiles) > 0 {
		buff.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\r\n", boundary))

		// 写入正文的边界信息
		buff.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	}

	// 写入正文头部
	if isHTMLBody {
		buff.WriteString(fmt.Sprintf("Content-Type: text/html; charset=\"utf-8\"\r\n"))
	} else {
		buff.WriteString(fmt.Sprintf("Content-Type: text/plain; charset=\"utf-8\"\r\n"))
	}
	buff.WriteString("\r\n")
	// 写入正文内容
	buff.WriteString(body)

	if len(attachFiles) > 0 {
		for _, file := range attachFiles {
			fileBytes, err := ioutil.ReadFile(file)
			if err != nil {
				return err
			}

			_, fileName := filepath.Split(file)

			// 写入文件信息
			buff.WriteString(fmt.Sprintf("\r\n\r\n--%s\r\n", boundary))
			buff.WriteString("Content-Type: application/octet-stream\r\n")
			buff.WriteString("Content-Transfer-Encoding: base64\r\n")
			buff.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n\r\n", fileName))

			// 写入文件数据
			b := make([]byte, base64.StdEncoding.EncodedLen(len(fileBytes)))
			base64.StdEncoding.Encode(b, fileBytes)
			buff.Write(b)
		}

		// 写入结束标识
		buff.WriteString(fmt.Sprintf("\r\n--%s--", boundary))
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	// 写入邮件数据
	_, err = w.Write(buff.Bytes())
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}
