package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/rosas99/monster/pkg/log"
	"github.com/vmihailenco/msgpack/v5"
	"sync"
	"sync/atomic"
	"time"
)

// Pusher will record analytics data to a redis back end as defined in the Config object.
type Pusher struct {
	store                      *redis.Client
	poolSize                   int
	recordsChan                chan any
	workerBufferSize           uint64
	recordsBufferFlushInterval uint64
	shouldStop                 uint32
	poolWg                     sync.WaitGroup
	channel                    string
}

const (
	recordsBufferForcedFlushInterval = 1 * time.Second
)

var pushers *Pusher

// NewPusher returns a new analytics instance.
func NewPusher(options *PusherOptions, store *redis.Client, channel string) *Pusher {
	ps := options.PoolSize
	recordsBufferSize := options.RecordsBufferSize
	workerBufferSize := recordsBufferSize / uint64(ps)
	log.Debugw("Analytics pool worker buffer size", workerBufferSize)

	recordsChan := make(chan any, recordsBufferSize)

	pushers = &Pusher{
		store:                      store,
		poolSize:                   ps,
		recordsChan:                recordsChan,
		workerBufferSize:           workerBufferSize,
		recordsBufferFlushInterval: options.FlushInterval,
		channel:                    channel,
	}

	return pushers
}

// GetPusher returns the existed analytics instance.
// Need to initialize `analytics` instance before calling GetAnalytics.
func GetPusher() *Pusher {
	return pushers
}

// Start the analytics service.
func (p *Pusher) Start() {
	// start worker pool
	atomic.SwapUint32(&p.shouldStop, 0)
	for i := 0; i < p.poolSize; i++ {
		p.poolWg.Add(1)
		go p.recordWorker()
	}
}

// Stop the analytics service.
func (p *Pusher) Stop() {
	// flag to stop sending records into channel
	atomic.SwapUint32(&p.shouldStop, 1)

	// close channel to stop workers
	close(p.recordsChan)

	// wait for all workers to be done
	p.poolWg.Wait()
}

// Record will store an AnalyticsRecord in Redis.
func (p *Pusher) Record(record any) error {
	// check if we should stop sending records 1st
	if atomic.LoadUint32(&p.shouldStop) > 0 {
		return nil
	}

	// just send record to channel consumed by pool of workers
	// leave all data crunching and Redis I/O work for pool workers
	p.recordsChan <- record
	return nil
}

func (p *Pusher) recordWorker() {
	defer p.poolWg.Done()

	// this is buffer to send one pipelined command to redis
	// use r.recordsBufferSize as cap to reduce slice re-allocations
	recordsBuffer := make([][]byte, 0, p.workerBufferSize)

	// read records from channel and process
	lastSentTS := time.Now()

	for {
		var readyToSend bool
		select {
		case record, ok := <-p.recordsChan:
			// check if channel was closed and it is time to exit from worker
			if !ok {
				// send what is left in buffer
				p.store.Publish(context.Background(), p.channel, recordsBuffer)
				return
			}

			// we have new record - prepare it and add to buffer
			//if encoded, err := json.Marshal(record); err != nil {
			if encoded, err := msgpack.Marshal(record); err != nil {
				log.Errorf("Error encoding analytics data: %s", err)
			} else {
				recordsBuffer = append(recordsBuffer, encoded)
			}

			// identify that buffer is ready to be sent
			readyToSend = uint64(len(recordsBuffer)) == p.workerBufferSize
		case <-time.After(time.Duration(p.recordsBufferFlushInterval) * time.Millisecond):
			// nothing was received for that period of time
			// anyway, send whatever we have, don't hold data too long in buffer
			readyToSend = true
		}

		// send data to Redis and reset buffer
		if len(recordsBuffer) > 0 && (readyToSend || time.Since(lastSentTS) >= recordsBufferForcedFlushInterval) {
			p.store.Publish(context.Background(), p.channel, recordsBuffer)
			recordsBuffer = recordsBuffer[:0]
			lastSentTS = time.Now()
		}

	}
}

// DurationToMillisecond convert time duration type to float64.
func DurationToMillisecond(d time.Duration) float64 {
	return float64(d) / 1e6
}
