// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rosas99/monster/internal/sms/model"
	"github.com/rosas99/monster/internal/sms/store"
	"github.com/rosas99/monster/pkg/log"
	genericoptions "github.com/rosas99/monster/pkg/options"
	"github.com/segmentio/kafka-go"
)

// kafkaLogger is a log.Logger implementation that writes log messages to Kafka.
type Logger struct {
	// enabled is an atomic boolean indicating whether the logger is enabled.
	enabled int32
	// writer is the Kafka writer used to write log messages.

	writer *kafka.Writer
	ds     store.HistoryStore
}

// todo 这里改成短信发送历史
// AuditMessage is the message structure for log messages.
type AuditMessage struct {
	Matcher   string     `protobuf:"bytes,1,opt,name=matcher,proto3" json:"matcher,omitempty"`
	Request   []any      `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Result    bool       `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	Explains  [][]string `protobuf:"bytes,4,opt,name=explains,proto3" json:"explains,omitempty"`
	Timestamp int64      `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

// NewLogger creates a new kafkaLogger instance.
func NewLogger(kafkaOpts *genericoptions.KafkaOptions, ds store.HistoryStore) (*Logger, error) {
	writer, err := kafkaOpts.Writer()
	if err != nil {
		return nil, err
	}

	return &Logger{writer: writer, ds: ds}, nil
}

// LogModel writes a log message for the policy model.
func (l *Logger) LogHistory(history *model.HistoryM) {
	err := l.ds.Create(context.Background(), history)
	if err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	}
}

// log others

// LogModel writes a log message for the policy model.
func (l *Logger) LogKpi(messageMap map[string]any) {

	out, _ := json.Marshal(messageMap)

	fmt.Println(messageMap)
	if err := l.writer.WriteMessages(context.Background(), kafka.Message{Value: out}); err != nil {
		log.Errorw(err, "Failed to write kafka messages")
	} else {
		fmt.Println(string(out))
	}
}
