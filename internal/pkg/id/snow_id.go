package id

import (
	"context"
	"github.com/rosas99/monster/pkg/id"
)

func GenerateSnowId() uint64 {
	options := func(*id.SonyflakeOptions) {
		id.WithSonyflakeMachineId(1)
	}

	snowIns := id.NewSonyflake(options)
	return snowIns.Id(context.Background())
}
