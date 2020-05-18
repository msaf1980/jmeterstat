package statcalc

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatCalculator(t *testing.T) {
	var stat StatCalculator
	// Init StatCalc
	stat.Init()

	var tests = []struct {
		descr string
		input []float64
		reset bool
		isNaN bool
		count uint64
		min   float64 // expected result, not checked when NaN
		max   float64
		sum   float64
		mean  float64
		p90   float64
	}{
		{
			"{}",
			[]float64{}, false, true,
			0, //count
			0.0, 0.0, 0.0, 0.0,
			0.0, // p90
		},
		{
			"{1.0, 2.0, 3.0, 0.5}",
			[]float64{1.0, 2.0, 3.0, 0.5}, false, false,
			4, //count
			0.5, 3.0, 6.5, 6.5 / 4,
			2.5, // p90
		},
		{
			"{1.0, 2.0, 3.0, 0.5, 0.0}",
			[]float64{0.0}, false, false,
			5, //count
			0.0, 3.0, 6.5, 6.5 / 5,
			2.5, // p90
		},
		{
			"{1.0, 2.0, 3.0, 0.5, 0.0, 1.0, 4.0, 2.0, 0.5, 0.0}",
			[]float64{1.0, 4.0, 2.0, 0.5, 0.0}, false, false,
			10, //count
			0.0, 4.0, 6.5 + 1.0 + 4.0 + 2.0 + 0.5, (6.5 + 1.0 + 4.0 + 2.0 + 0.5) / 10.0,
			3.0, // p90
		},
		{
			"{1.0, 2.0, 3.0, 0.5, 0.0, 1.0, 4.0, 2.0, 0.5, 0.0, 0.0}",
			[]float64{0.0}, false, false,
			11, //count
			0.0, 4.0, 6.5 + 1.0 + 4.0 + 2.0 + 0.5, (6.5 + 1.0 + 4.0 + 2.0 + 0.5) / 11.0,
			2.5, // p90
		},
		{
			"{1.0}",
			[]float64{1.0}, true, false,
			1, //count
			1.0, 1.0, 1.0, 1.0,
			1.0, // p90
		},
	}

	for _, tt := range tests {
		if tt.reset {
			stat.Init()
		}

		for _, v := range tt.input {
			stat.AddValue(v)
		}
		assert.Equalf(t, tt.count, stat.Count(), "%s.Count() failed", tt.descr)
		if tt.isNaN {
			assert.True(t, math.IsNaN(stat.Min()), "%s.Min() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Max()), "%s.Max() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Sum()), "%s.Sum() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Mean()), "%s.Mean() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Percentile(90)), "%s.Percentile(90) failed", tt.descr)
		} else {
			assert.Equalf(t, tt.min, stat.Min(), "%s.Min() failed", tt.descr)
			assert.Equalf(t, tt.max, stat.Max(), "%s.Max() failed", tt.descr)
			assert.Equalf(t, tt.sum, stat.Sum(), "%s.Sum() failed", tt.descr)
			assert.Equalf(t, tt.mean, stat.Mean(), "%s.Mean() failed", tt.descr)
			assert.Equalf(t, tt.p90, stat.Percentile(90), "%s.Percentile(90) failed", tt.descr)
		}
	}
}

func TestStatCalculatorWithDefault(t *testing.T) {
	var stat StatCalculator
	// Init StatCalc
	stat.InitWithDefaults(0, 0, 0)

	var tests = []struct {
		descr string
		input []float64
		reset bool
		isNaN bool
		count uint64
		min   float64 // expected result, not checked when NaN
		max   float64
		sum   float64
		mean  float64
		p90   float64
	}{
		{
			"{}",
			[]float64{}, false, false,
			0, //count
			0.0, 0.0, 0.0, 0.0,
			0.0, // p90
		},
		{
			"{1.0, 2.0, 3.0, 0.5}",
			[]float64{1.0, 2.0, 3.0, 0.5}, false, false,
			4, //count
			0.5, 3.0, 6.5, 6.5 / 4,
			2.5, // p90
		},
		{
			"{1.0}",
			[]float64{1.0}, true, false,
			1, //count
			1.0, 1.0, 1.0, 1.0,
			1.0, // p90
		},
	}

	for _, tt := range tests {
		if tt.reset {
			stat.Init()
		}

		for _, v := range tt.input {
			stat.AddValue(v)
		}
		assert.Equalf(t, tt.count, stat.Count(), "%s.Count() failed", tt.descr)
		if tt.isNaN {
			assert.True(t, math.IsNaN(stat.Min()), "%s.Min() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Max()), "%s.Max() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Sum()), "%s.Sum() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Mean()), "%s.Mean() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Percentile(90)), "%s.Percentile(90) failed", tt.descr)
		} else {
			assert.Equalf(t, tt.min, stat.Min(), "%s.Min() failed", tt.descr)
			assert.Equalf(t, tt.max, stat.Max(), "%s.Max() failed", tt.descr)
			assert.Equalf(t, tt.sum, stat.Sum(), "%s.Sum() failed", tt.descr)
			assert.Equalf(t, tt.mean, stat.Mean(), "%s.Mean() failed", tt.descr)
			assert.Equalf(t, tt.p90, stat.Percentile(90), "%s.Percentile(90) failed", tt.descr)
		}
	}
}

func TestStatCalculatorWithNonZero(t *testing.T) {
	var stat StatCalculator
	// Init StatCalc
	stat.InitWithDefaults(-1.0, -1.0, 0.0)

	var tests = []struct {
		descr string
		input []float64
		reset bool
		isNaN bool
		count uint64
		min   float64 // expected result, not checked when NaN
		max   float64
		sum   float64
		mean  float64
		p90   float64
	}{
		{
			"{}",
			[]float64{}, false, false,
			0, //count
			-1.0, 0.0, 0.0, 0.0,
			-1.0, // p90
		},
		{
			"{1.0, 2.0, 3.0, 0.5}",
			[]float64{1.0, 2.0, 3.0, 0.5}, false, false,
			4, //count
			0.5, 3.0, 6.5, 6.5 / 4,
			2.5, // p90
		},
		{
			"{1.0}",
			[]float64{1.0}, true, false,
			1, //count
			1.0, 1.0, 1.0, 1.0,
			1.0, // p90
		},
	}

	for _, tt := range tests {
		if tt.reset {
			stat.Init()
		}

		for _, v := range tt.input {
			stat.AddValue(v)
		}
		assert.Equalf(t, tt.count, stat.Count(), "%s.Count() failed", tt.descr)
		if tt.isNaN {
			assert.True(t, math.IsNaN(stat.Min()), "%s.Min() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Max()), "%s.Max() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Sum()), "%s.Sum() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Mean()), "%s.Mean() failed", tt.descr)
			assert.True(t, math.IsNaN(stat.Percentile(90)), "%s.Percentile(90) failed", tt.descr)
		} else {
			assert.Equalf(t, tt.min, stat.Min(), "%s.Min() failed", tt.descr)
			assert.Equalf(t, tt.max, stat.Max(), "%s.Max() failed", tt.descr)
			assert.Equalf(t, tt.sum, stat.Sum(), "%s.Sum() failed", tt.descr)
			assert.Equalf(t, tt.mean, stat.Mean(), "%s.Mean() failed", tt.descr)
			assert.Equalf(t, tt.p90, stat.Percentile(90), "%s.Percentile(90) failed", tt.descr)
		}
	}
}

func Test_compareWithNaN(t *testing.T) {
	type args struct {
		s float64
		o float64
	}
	tests := []struct {
		name  string
		args  args
		equal bool
	}{
		{
			"NaN, NaN",
			args{math.NaN(), math.NaN()},
			true,
		},
		{
			"NaN, 0",
			args{math.NaN(), 0},
			false,
		},
		{
			"1, NaN",
			args{1, math.NaN()},
			false,
		},
		{
			"0, 0",
			args{0, 0},
			true,
		},
		{
			"1, 0",
			args{1, 0},
			false,
		},
	}
	for _, tt := range tests {
		if got := compareWithNaN(tt.args.s, tt.args.o); got != tt.equal {
			t.Errorf("compareWithNaN(%s) = %v, want %v", tt.name, got, tt.equal)
		}
	}
}

func TestStatCalculator_Equal(t *testing.T) {
	type args struct {
		s *StatCalculator
		o *StatCalculator
	}
	tests := []struct {
		name  string
		args  args
		equal bool
	}{
		{
			"{}() == {}()",
			args{
				new(StatCalculator).Init(),
				new(StatCalculator).Init(),
			},
			true,
		},
		{
			"{}(0.0, 0.0, 0.0) == {}(0.0, 0.0, 0.0)",
			args{
				new(StatCalculator).InitWithDefaults(0.0, 0.0, 0.0),
				new(StatCalculator).InitWithDefaults(0.0, 0.0, 0.0),
			},
			true,
		},
		{
			"{}() == {}(0.0, 0.0, 0.0)",
			args{
				new(StatCalculator).Init(),
				new(StatCalculator).InitWithDefaults(0.0, 0.0, 0.0),
			},
			false,
		},
		{
			"{}(0.0, 0.0, 0.0) == {}()",
			args{
				new(StatCalculator).InitWithDefaults(0.0, 0.0, 0.0),
				new(StatCalculator).Init(),
			},
			false,
		},
		{
			"{1.0} == {1.0}",
			args{
				new(StatCalculator).Init().AddValue(1.0),
				new(StatCalculator).Init().AddValue(1.0),
			},
			true,
		},
		{
			"{1.0, 0.1} == {1.0, 0.1}",
			args{
				new(StatCalculator).Init().AddValue(1.0).AddValue(0.1),
				new(StatCalculator).Init().AddValue(1.0).AddValue(0.1),
			},
			true,
		},
		{
			"{1.0} == {1.0, 0.1}",
			args{
				new(StatCalculator).Init().AddValue(1.0),
				new(StatCalculator).Init().AddValue(1.0).AddValue(0.1),
			},
			false,
		},
		{
			"{1.0, 0.1} == {1.0}",
			args{
				new(StatCalculator).Init().AddValue(1.0).AddValue(0.1),
				new(StatCalculator).Init().AddValue(1.0),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.s.Equal(tt.args.o); got != tt.equal {
				t.Errorf("(%s) = %v, want %v", tt.name, got, tt.equal)
			}
		})
	}
}
