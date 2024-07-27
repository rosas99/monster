package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/rosas99/monster/internal/fakeserver/biz IBiz

import (
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/usercenter/biz/user"
	"github.com/rosas99/monster/internal/usercenter/store"
	"github.com/segmentio/kafka-go"
)

// IBiz defines the methods that need to be implemented in the Biz layer.
type IBiz interface {
	Users() user.IBiz
}

// Biz is a concrete implementation of the IBiz interface.
type Biz struct {
	ds          store.IStore
	rds         *redis.Client
	kafkaWriter *kafka.Writer
}

// Ensure that biz implements the IBiz interface.
var _ IBiz = (*Biz)(nil)

// NewBiz creates an instance of type IBiz.
func NewBiz(ds store.IStore, rds *redis.Client) *Biz {
	return &Biz{ds: ds, rds: rds}
}

// Users returns an instance that implements the IBiz interface.
func (b *Biz) Users() user.IBiz {
	return user.New(b.ds, b.rds)
}
