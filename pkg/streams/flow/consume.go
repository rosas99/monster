// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas99/monster.
//

package flow

import (
	"github.com/rosas99/monster/pkg/streams"
)

type ConsumeHandler interface {
	Consume(val any) error
}

// ConsumeHandler represents a Consumer transformation interface.

// Consumer takes one element and produces one element.
//
// in  -- 1 -- 2 ---- 3 -- 4 ------ 5 --
//
// [ ---------- ConsumeHandler ---------- ]
//
// out -- 1' - 2' --- 3' - 4' ----- 5' -.
type Consumer struct {
	in          chan any
	out         chan any
	parallelism uint
	handler     ConsumeHandler
}

// Verify Map satisfies the Flow interface.
//var _ streams.Flow = (*Map[any, any])(nil)

// NewConsumer NewMap returns a new Map instance.
//
// mapFunction is the Map transformation function.
// parallelism is the flow parallelism factor. In case the events order matters, use parallelism = 1.
func NewConsumer(handler ConsumeHandler, parallelism uint) *Consumer {
	mapFlow := &Consumer{
		handler:     handler,
		in:          make(chan any),
		out:         make(chan any),
		parallelism: parallelism,
	}
	go mapFlow.doStream()
	return mapFlow
}

// Via streams data through the given flow.
func (m *Consumer) Via(flow streams.Flow) streams.Flow {
	go m.transmit(flow)
	return flow
}

// To streams data to the given sink.
func (m *Consumer) To(sink streams.Sink) {
	m.transmit(sink)
}

// Out returns an output channel for sending data.
func (m *Consumer) Out() <-chan any {
	return m.out
}

// In returns an input channel for receiving data.
func (m *Consumer) In() chan<- any {
	return m.in
}

func (m *Consumer) transmit(inlet streams.Inlet) {
	for element := range m.Out() {
		inlet.In() <- element
	}
	close(inlet.In())
}

// 执行函数的方法
func (m *Consumer) doStream() {
	sem := make(chan struct{}, m.parallelism)
	for elem := range m.in {
		sem <- struct{}{}
		go func(elem any) {
			defer func() { <-sem }()
			if err := m.handler.Consume(elem); err != nil {
				// 处理错误
				// 提交消息等处理

			}
			m.out <- elem
		}(elem)
	}
	for i := 0; i < int(m.parallelism); i++ {
		sem <- struct{}{}
	}
	close(m.out)
}
