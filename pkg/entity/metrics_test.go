package entity

import "testing"

func TestMetricEqual(t *testing.T) {
	type args struct {
		l1 []Label
		l2 []Label
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "len 1 equal",
			args: args{
				l1: []Label{
					{Name: "n1", Value: "val1"},
				},
				l2: []Label{
					{Name: "n1", Value: "val1"},
				},
			},
			want: true,
		},
		{
			name: "len>1 equal(sorted)",
			args: args{
				l1: []Label{
					{Name: "n1", Value: "val1"},
					{Name: "n2", Value: "val2"},
				},
				l2: []Label{
					{Name: "n1", Value: "val1"},
					{Name: "n2", Value: "val2"},
				},
			},
			want: true,
		},
		{
			name: "len>1 equal(unsorted)",
			args: args{
				l1: []Label{
					{Name: "n1", Value: "val1"},
					{Name: "n2", Value: "val2"},
				},
				l2: []Label{
					{Name: "n2", Value: "val2"},
					{Name: "n1", Value: "val1"},
				},
			},
			want: true,
		},
		{
			name: "len not equal",
			args: args{
				l1: []Label{
					{Name: "n1", Value: "val1"},
				},
				l2: []Label{
					{Name: "n1", Value: "val1"},
					{Name: "n2", Value: "val2"},
				},
			},
			want: false,
		},
		{
			name: "len=1 name not equal",
			args: args{
				l1: []Label{
					{Name: "n1", Value: "val1"},
				},
				l2: []Label{
					{Name: "n2", Value: "val1"},
				},
			},
			want: false,
		},
		{
			name: "len>1 not equal",
			args: args{
				l1: []Label{
					{Name: "n1", Value: "val1"},
					{Name: "n2", Value: "val2"},
				},
				l2: []Label{
					{Name: "n1", Value: "val1"},
					{Name: "n2", Value: "val3"},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SeriesEqual(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("MetricEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
