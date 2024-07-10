package main

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		core.WriteResponse(c, nil, gin.H{
//			"code": 0,
//			"msg":  "ok",
//		})
//
//	})
//	r.Run("127.0.0.1:9494") // 监听并在 0.0.0.0:8080 上启动服务
//}

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	time "github.com/alibabacloud-go/darabonba-time/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 使用AK&SK初始化账号Client
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{}
	config.AccessKeyId = accessKeyId
	config.AccessKeySecret = accessKeySecret
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}

func _main() (_err error) {
	client, _err := CreateClient(tea.String("LTAI5t8Kd2TFFgzNWukJuz3e"), tea.String("TVQ2SkfObARCxw3bmW7eba71HBP2oN"))
	if _err != nil {
		return _err
	}

	// 1.发送短信
	numbers := "17806705418"
	signName := "Mervyn"
	templateCode := "SMS_464076579"
	templateParam := "{\"code\":\"1234\"}"
	sendReq := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  &numbers,
		SignName:      &signName,
		TemplateCode:  &templateCode,
		TemplateParam: &templateParam,
	}
	sendResp, _err := client.SendSms(sendReq)
	if _err != nil {
		return _err
	}

	code := sendResp.Body.Code
	if !tea.BoolValue(util.EqualString(code, tea.String("OK"))) {
		console.Log(tea.String("错误信息: " + tea.StringValue(sendResp.Body.Message)))
		return _err
	}

	bizId := sendResp.Body.BizId
	// 2. 等待 10 秒后查询结果
	_err = util.Sleep(tea.Int(10000))
	if _err != nil {
		return _err
	}
	// 3.查询结果
	phoneNums := string_.Split(&numbers, tea.String(","), tea.Int(-1))
	for _, phoneNum := range phoneNums {
		queryReq := &dysmsapi.QuerySendDetailsRequest{
			PhoneNumber: util.AssertAsString(phoneNum),
			BizId:       bizId,
			SendDate:    time.Format(tea.String("yyyyMMdd")),
			PageSize:    tea.Int64(10),
			CurrentPage: tea.Int64(1),
		}
		queryResp, _err := client.QuerySendDetails(queryReq)
		if _err != nil {
			return _err
		}

		dtos := queryResp.Body.SmsSendDetailDTOs.SmsSendDetailDTO
		// 打印结果
		for _, dto := range dtos {
			if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("3"))) {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送成功，接收时间: " + tea.StringValue(dto.ReceiveDate)))
			} else if tea.BoolValue(util.EqualString(tea.String(tea.ToString(tea.Int64Value(dto.SendStatus))), tea.String("2"))) {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 发送失败"))
			} else {
				console.Log(tea.String(tea.StringValue(dto.PhoneNum) + " 正在发送中..."))
			}

		}
	}
	return _err
}

func main() {
	err := _main()
	if err != nil {
		panic(err)
	}
}
