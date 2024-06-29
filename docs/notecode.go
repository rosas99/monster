package docs

import (
	"context"
	"fmt"
	"github.com/superproj/onex/pkg/id"
)

func main() {
	// 整个可以设置为全局变量 只初始化一次
	options := func(*id.SonyflakeOptions) {
		id.WithSonyflakeMachineId(1) // 自定义机器ID，默认为自动检测
	}

	snowIns := id.NewSonyflake(options)
	id := snowIns.Id(context.Background())
	fmt.Print("id is :", id)
}
