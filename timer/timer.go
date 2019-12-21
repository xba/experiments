package timer

import (
	"time"
)

const (
	sizeDays    = 100
	maskDays    = sizeDays - 1
	sizeHours   = 24
	maskHours   = sizeHours - 1
	sizeMinutes = 60
	maskMinutes = sizeMinutes - 1
	sizeSeconds = 60
	maskSeconds = sizeSeconds - 1
)

type Item uint64

type node struct {
	item Item
	ttl  [4]uint8
}

type Wheel struct {
	now      time.Time
	buckets  [][][]node
	onExpire func(items ...Item)
}

func NewWheel(onExpire func(items ...Item)) *Wheel {
	wheel := &Wheel{
		onExpire: onExpire,
		buckets: [][][]node{
			make([][]node, sizeDays),
			make([][]node, sizeHours),
			make([][]node, sizeMinutes),
			make([][]node, sizeSeconds),
		},
	}
	return wheel
}

// TODO
func (w *Wheel) Tick() {}

func (w *Wheel) Add(ttl time.Duration, items ...Item) {
	d := parse(ttl)
	// round down days to sizeDays
	//
	// TODO: in the future, might just want to return an error here so the user
	//       is forced to acknowledge the ttl limit
	if d[0] > sizeDays {
		d[0] = sizeDays
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
			if *bucket == nil {
				*bucket = nodes
			} else {
				*bucket = append(*bucket, nodes...)
			}
			return
		}
	}
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
