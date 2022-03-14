package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Sms struct {
	Account  string
	Password string
	Host     string
}

// RegistrySms 定义接口，在接口里面定义要实现的方法
type RegistrySms interface {
	// SendSms 定义发送短信的方法
	SendSms(mobile, template string) error
}

func NewRegistrySms(account, password, host string) RegistrySms {
	return &Sms{
		Account:  account,
		Password: password,
		Host:     host,
	}
}

// SendSms 发送手机短信
func (s Sms) SendSms(mobile, template string) error {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = s.Account   //创蓝API账号
	params["password"] = s.Password //创蓝API密码
	params["phone"] = mobile        //手机号码

	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	params["msg"] = url.QueryEscape(template)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New(err.Error())
	}
	reader := bytes.NewReader(bytesData)
	//短信发送URL
	smsUrl := fmt.Sprintf("http://%s/%s", s.Host, "/msg/send/json")
	request, err := http.NewRequest("POST", smsUrl, reader)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New(err.Error())
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New(err.Error())
	}
	var smsResponse = make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&smsResponse)
	if err != nil {
		fmt.Println("转码失败")
		return errors.New(err.Error())
	}
	fmt.Println("发送短信结果", smsResponse)
	if smsResponse["code"] == "0" {
		return nil
	} else {
		// 接口数据类型断言成字符类型
		message := smsResponse["errorMsg"]
		result, _ := message.(string)
		return errors.New(result)
	}
}
