package timer

import (
	"fmt"
	"sync"
	"time"
)

type Item uint64

type node struct {
	item Item
	ttl  [4]uint8
}

type bucket struct {
	sync.RWMutex
	nodes []node
}

type Wheel struct {
	last     time.Time
	cursor   uint64
	buckets  [][]bucket
	onExpire func(items ...Item)
	stop     chan struct{}
}

func NewWheel(onExpire func(items ...Item)) *Wheel {
	wheel := &Wheel{
		last:     time.Now(),
		onExpire: onExpire,
		buckets: [][]bucket{
			make([]bucket, 100),
			make([]bucket, 24),
			make([]bucket, 60),
			make([]bucket, 60),
		},
		stop: make(chan struct{}),
	}
	return wheel
}

func (w *Wheel) Run() {
	// tick every 250ms
	ticker := time.NewTicker(time.Second / 5)
	defer ticker.Stop()
	for {
		select {
		case <-w.stop:
			return
		case t := <-ticker.C:
			if t.Sub(w.last) > time.Second {
				w.last = t
				w.tick()
			}
		}
	}
}

func (w *Wheel) Add(ttl time.Duration, items ...Item) {
	if items == nil || w == nil {
		return
	}
	// extract days, hours, minutes, and seconds from the ttl duration
	d := parse(ttl)
	// round down days to 100
	//
	// TODO: in the future, might just want to return an error here so the user
	//       is forced to acknowledge the ttl limit
	if d[0] > 100 {
		d[0] = 100
	}
	// creates node(s)
	nodes := make([]node, len(items))
	for i := range nodes {
		nodes[i] = node{item: items[i], ttl: d}
	}
	// place node(s) in the highest bucket
	for i := range d {
		if d[i] != 0 {
			bucket := &w.buckets[i][d[i]]
			bucket.Lock()
			if bucket.nodes == nil {
				bucket.nodes = nodes
			} else {
				bucket.nodes = append(bucket.nodes, nodes...)
			}
			bucket.Unlock()
			fmt.Println(bucket.nodes)
			return
		}
	}
}

func (w *Wheel) tick() {
	w.cursor++
}

// parse takes a nanosecond duration and returns the respective days, hours
// minutes, and seconds
func parse(n time.Duration) (t [4]uint8) {
	// days
	t[0] = uint8(n / 8.64e13)
	// hours
	t[1] = uint8(n / 3.6e12 % 24)
	// minutes
	t[2] = uint8(n / 6e10 % 60)
	// seconds
	t[3] = uint8(n / 1e9 % 60)
	return
}

func parseTicks(n uint64) (t [4]uint8) {
	// days
	t[0] = uint8(n / 86400)
	// hours
	t[1] = uint8(n / 3600 % 24)
	// minutes
	t[2] = uint8(n / 60 % 60)
	// seconds
	t[3] = uint8(n % 60)
	return
}
