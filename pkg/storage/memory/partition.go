package memory

import "time"

type Partition struct {
	MinTimestamp int64
	// MaxTimestamp int64

	Body []Slot
}

func NewPartition() Partition {
	min := time.Now().UnixMilli()
	// max := min + cfg.PartitionDuration.Milliseconds()
	return Partition{
		MinTimestamp: min,
		// MaxTimestamp: max,
		Body: make([]Slot, 0),
	}
}
