package id

import (
	"context"
	"github.com/rosas99/monster/pkg/id"
)

func GenerateId() uint64 {
	options := func(*id.SonyflakeOptions) {
		id.WithSonyflakeMachineId(1) // 自定义机器ID，默认为自动检测
	}

	snowIns := id.NewSonyflake(options)
	return snowIns.Id(context.Background())
}
