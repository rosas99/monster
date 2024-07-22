package biz

//go:generate mockgen -destination mock_biz.go -package biz github.com/rosas99/monster/internal/fakeserver/biz IBiz

import (
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/internal/pkg/idempotent"
	"github.com/rosas99/monster/internal/sms/biz/interaction"
	"github.com/rosas99/monster/internal/sms/biz/message"
	"github.com/rosas99/monster/internal/sms/biz/template"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/internal/sms/writer"
	"github.com/segmentio/kafka-go"
)

// IBiz defines a set of methods for returning interfaces that the biz struct implements.
type IBiz interface {
	Templates() template.IBiz
	Messages() message.IBiz
	Interaction() interaction.IBiz
}

type biz struct {
	ds          store.IStore
	rds         *redis.Client
	idt         *idempotent.Idempotent
	logger      *writer.Writer
	kafkaWriter *kafka.Writer
}

// Ensure biz implements IBiz.
var _ IBiz = (*biz)(nil)

// NewBiz returns a pointer to a new instance of the biz struct.
func NewBiz(ds store.IStore, rds *redis.Client, idt *idempotent.Idempotent, logger *writer.Writer) *biz {
	return &biz{ds: ds, rds: rds, idt: idt, logger: logger}
}

// Templates returns a new instance of the IBiz interface.
func (b *biz) Templates() template.IBiz {
	return template.New(b.ds, b.rds)
}

// Messages returns a new instance of the IBiz interface.
func (b *biz) Messages() message.IBiz {
	return message.New(b.ds, b.logger, b.rds, b.idt)
}

// Interaction returns a new instance of the IBiz interface.
func (b *biz) Interaction() interaction.IBiz {
	return interaction.New(b.ds, b.logger, b.rds, b.idt)
}
