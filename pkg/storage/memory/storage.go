package memory

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/x-carno/carnots/pkg/entity"
	"github.com/x-carno/carnots/pkg/storage/db"
)

var mstorage *Storage

type Storage struct {
	// Partitions []Partition

	Blocks []Block

	sync.Mutex
}

func NewStorage() db.DB {
	if mstorage != nil {
		return mstorage
	}
	mstorage = new(Storage)
	return mstorage
}

// TODO : use block slice when appending metrics for efficiency, retain the slice for a Partition Duration
// compact metrics(need to think about a compact algorithm) every Partition Duration to reduce RAM occupation

// usually, there is just one batch of metrics at one time, that means len(metrics)=1
func (s *Storage) Store(metrics ...entity.Metrics) {
	if len(metrics) == 0 {
		logrus.Warnln("no metrics passed in")
		return
	}

	for _, metric := range metrics {
		s.AppendMetric(metric)
	}
}

func (s *Storage) AppendMetric(metric entity.Metrics) {
	for _, v := range metric.Series {
		s.AppendSeries(v)
	}
}

func (s *Storage) AppendSeries(series entity.TimeSeries) {
	hasLabels := false
	for _, block := range s.Blocks {
		if entity.SeriesEqual(block.Labels, series.Labels) {
			hasLabels = true
			block.AppendSamples(series.Samples)
		}
	}
	if !hasLabels {
		block := NewBlock(series.Labels)
		block.AppendSamples(series.Samples)
		s.Blocks = append(s.Blocks, block)
	}
}

// func (s *Storage) FindOrCreate(labels []storage.Label) Block {
// 	for _, block := range s.Blocks {
// 		if storage.SeriesEqual(block.Labels, labels) {
// 			return block
// 		}
// 	}
// 	return NewBlock(labels)
// }
