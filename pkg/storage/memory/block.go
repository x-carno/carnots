package memory

import (
	"sort"
	"time"

	"github.com/x-carno/carnots/pkg/entity"
)

// block is append-only
type Block struct {
	// BlockMetrics []BlockMetric

	// the value of label '__name__' is the metric's name
	Labels []entity.Label

	Pts []Partition
}

type BlockMetric struct {
	// the value of label '__name__' is the metric's name
	Labels []entity.Label

	Partitions []Partition
}

func NewBlock(labels []entity.Label) Block {
	// sort labels according to label name
	sort.Slice(labels, func(i, j int) bool {
		return labels[i].Name < labels[j].Name
	})
	return Block{
		Labels: labels,
		Pts:    []Partition{NewPartition()},
	}
}

func (b *Block) AppendSamples(samples []entity.Sample) {
	for i := 0; i < len(samples); i++ {
		sample := samples[i]
		if sample.Timestamp == 0 {
			sample.Timestamp = time.Now().UnixMilli() + int64(i)*cfg.Precision.Milliseconds()
		}
		ptsIdx := len(b.Pts) - 1
		if sample.Timestamp <= b.Pts[ptsIdx].MinTimestamp+cfg.PartitionDuration.Milliseconds() {
			b.Pts[ptsIdx].Body = append(b.Pts[ptsIdx].Body, Slot{
				Timestamp: sample.Timestamp,
				Value:     sample.Value,
			})
		} else {
			newPts := NewPartition()
			newPts.Body = append(newPts.Body, Slot{
				Timestamp: sample.Timestamp,
				Value:     sample.Value,
			})
			b.Pts = append(b.Pts, newPts)
		}
	}
}
