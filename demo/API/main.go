package main
import (
    "net/http"
    "net/url"
    "encoding/json"
    "fmt"
    "bytes"
    "io/ioutil"
    "unsafe"
)
 
type JsonPostSample struct {
 
}
 
func  main() {
    params := make(map[string]interface{})
	
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
    params["account"] = ""  //创蓝API账号 
    params["password"] = "" //创蓝API密码
    params["phone"] = "18721755342" //手机号码
	
	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
    params["msg"] =url.QueryEscape("【253云通讯】您好，您的验证码是999999")  
    params["report"] = "true"
    bytesData, err := json.Marshal(params)
    if err != nil {
        fmt.Println(err.Error() )
        return
    }
    reader := bytes.NewReader(bytesData)
    url := "http://xxx/msg/send/json"  //短信发送URL 
    request, err := http.NewRequest("POST", url, reader)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    request.Header.Set("Content-Type", "application/json;charset=UTF-8")
    client := http.Client{}
    resp, err := client.Do(request)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    respBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
   
    str := (*string)(unsafe.Pointer(&respBytes))
    fmt.Println(*str)
}