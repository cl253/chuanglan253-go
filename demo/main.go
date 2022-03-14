package main

import (
    "fmt"
    "demo/sms"
)

func  main() {
    var account = ""  // 创蓝API账号 
    var password = "" // 创蓝API密码
    var host = "smssh1.253.com" // 创蓝的host一般直接写死就可以
    smsRegistry := sms.NewRegistrySms(account, password, host)
    // 发送短信
    if err := smsRegistry.SendSms("181xxxxx","【253云通讯】您好，您的验证码是999999");err!=nil {
        fmt.Printf("发送短信失败")
    }
}