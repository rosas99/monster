// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://"github.com/rosas99/monster.
//

package v1 // import ""github.com/rosas99/monster/pkg/api/fakeserver/v1"

// kratos proto client api/sms/v1/sms.proto

/*
 linux
protoc --proto_path=. \
       --proto_path=./third_party \
       --go_out=paths=source_relative:. \
       --go-errors_out=paths=source_relative:. \
       api/sms/errors.proto
  windows 只需把 \ 替换为 ^
protoc --proto_path=. --proto_path=./third_party --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. api\sms\errors.proto

// todo 确认路径怎么写

*/
