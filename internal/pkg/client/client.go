// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://"github.com/rosas99/monster.
//

package client

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/pflag"
)

// Define global options for all clients.
var (
	UserAgent  = "onex"
	Debug      = false
	RetryCount = 3
	Timeout    = 30 * time.Second
)

func AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&UserAgent, "client.user-agent", UserAgent, ""+
		"Used to specify the Resty client User-Agent.")

	fs.BoolVar(&Debug, "client.debug", Debug, ""+
		"Enables the debug mode on Resty client.")

	fs.IntVar(&RetryCount, "client.retry-count", RetryCount, ""+
		"Enables retry on Resty client and allows you to set no. of retry count. Resty uses a Backoff mechanism.")

	fs.DurationVar(&Timeout, "client.timeout", Timeout, ""+
		"Request timeout for client.")
}

func NewRequest() *resty.Request {
	return resty.New().
		SetRetryCount(RetryCount).
		SetDebug(Debug).
		R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   UserAgent,
		})
}

// IsDiscoveryEndpoint used to determine if the given endpoint is a service discovery endpoint.
func IsDiscoveryEndpoint(server string) bool {
	return strings.HasPrefix(server, "discovery:///")
}

type AuthSuccess struct {
}

func main() {
	url := "http://example.com/api"
	request := NewRequest()
	response, err := request.
		// 这里换成结构体
		SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
		SetResult(&AuthSuccess{}).
		Post(url)
	if err != nil {

	}

	fmt.Print(response)

	if response.StatusCode() >= 400 {
		fmt.Printf("服务器返回错误状态码: %d\n", response.StatusCode())
		// 根据需要处理不同的错误状态码
	}

	//resp, err := client.R().
	//	SetBody(Article{
	//		Tags: []string{"new tag1", "new tag2"},
	//	}).
	//	SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
	//	SetError(&errorStruct). // 注意这里传入了错误结构体的地址
	//	Patch("https://myapp.com/articles/1234")
	//
	//if err != nil {
	//	fmt.Printf("请求失败: %v\n", err)
	//} else if resp.IsError() { // 检查是否是错误响应
	//	fmt.Printf("错误码: %d\n", errorStruct.Code)
	//	// 根据需要处理错误码
	//}

	//例如，如果你的错误结构体Error有一个字段Code来存储错误码，你可以这样获取：
}
