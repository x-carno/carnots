package entity

import "sort"

type Metrics struct {
	Series []TimeSeries
}

func (m *Metrics) GetMinTimestamp() int64 {
	sort.Slice(m.Series, func(i, j int) bool {
		return m.Series[i].Samples[0].Timestamp < m.Series[j].Samples[0].Timestamp
	})
	return m.Series[0].Samples[0].Timestamp
}

type TimeSeries struct {
	Labels  []Label
	Samples []Sample
}

func (m *TimeSeries) Name() string {
	// var name string
	for _, label := range m.Labels {
		if label.Name == "__name__" {
			return label.Value
		}
	}
	return ""
}

// the value of label '__name__' is the metric's name
type Label struct {
	Name  string
	Value string
}

type Sample struct {
	Value float64

	// timestamp in milliseconds
	Timestamp int64
}

func SeriesEqual(l1, l2 []Label) bool {
	len1 := len(l1)
	len2 := len(l2)
	if len1 != len2 {
		return false
	}
	// sort.Slice(l1, func(i, j int) bool {
	// 	return l1[i].Name < l1[j].Name
	// })
	// sort.Slice(l2, func(i, j int) bool {
	// 	return l2[i].Name < l2[j].Name
	// })

	// lables are already sorted by names when block created
	for i := 0; i < len1; i++ {
		if l1[i].Name != l2[i].Name {
			return false
		}
		if l1[i].Value != l2[i].Value {
			return false
		}
	}
	return true
}
