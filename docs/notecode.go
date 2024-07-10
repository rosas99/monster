package docs

import (
	"context"
	"fmt"
	"github.com/rosas99/monster/pkg/id"
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

// option demo
/*
type Option func(client *Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}*/

/*
	//sem := make(chan struct{}, ks.parallelism)
	//for elem := range ks.out {
	//	sem <- struct{}{}
	//	go func(elem any) {
	//		defer func() { <-sem }()
	//		if err := ks.handler.Consume(elem); err != nil {
	//			suc <- false
	//		} else {
	//			suc <- true
	//		}
	//
	//		ks.out <- elem
	//	}(elem)
	//}
	//for i := 0; i < int(ks.parallelism); i++ {
	//	sem <- struct{}{}
	//}
	//close(ks.out)
*/
