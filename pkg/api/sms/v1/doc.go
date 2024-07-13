package v1 // import ""github.com/rosas99/monster/pkg/api/fakeserver/v1"

// kratos proto client api/sms/v1/sms.proto
// kratos proto client api/usercenter/v1/usercenter.proto

/*
 linux
protoc --proto_path=. \
       --proto_path=./third_party \
       --go_out=paths=source_relative:. \
       --go-errors_out=paths=source_relative:. \
       api/sms/errors.proto
  windows 只需把 \ 替换为 ^
protoc --proto_path=. --proto_path=./third_party --go_out=paths=source_relative:.
protoc --proto_path=. --proto_path=./third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.

// todo 确认路径怎么写

*/
