package statcalc

// Based on JMeter StatCalculator, but more primite and less functional
//   - simple flat array with values for percentiles
//   - no StandardDeviation, Distribution
//   - no bytes and sentBytes counters (I planned use additional instance of StatCalculator for detailed statistic)
// https://github.com/apache/jmeter/blob/b6d11d79d905d0c099732bb928d2372fd1388981/src/jorphan/src/main/java/org/apache/jorphan/math/StatCalculator.java
// At now not realized addEachValue(T val, long sampleCount) and addValue(T val, long sampleCount)

import (
	"math"

	"github.com/montanaflynn/stats"
)

const valuesInitSize = 100

// StatCalculator calculate stat (min/max/mean/percentiles)
type StatCalculator struct {
	sum float64

	min float64
	max float64

	mean float64

	values []float64

	count uint64

	zero float64
}

// Init reset StatCalculator to initial empthy state. Also call for next reuse.
func (s *StatCalculator) Init() *StatCalculator {
	s.zero = math.NaN()

	s.sum = math.NaN()
	s.min = math.NaN()
	s.max = math.NaN()
	s.mean = math.NaN()

	s.count = 0

	s.values = make([]float64, 0, valuesInitSize)

	return s
}

// InitWithDefaults reset StatCalculator to initial empthy state. But set sum to 0,
func (s *StatCalculator) InitWithDefaults(zero float64, min float64, max float64) *StatCalculator {
	s.zero = zero

	s.sum = 0
	s.min = min
	s.max = max
	s.mean = 0

	s.count = 0

	s.values = make([]float64, 0, valuesInitSize)

	return s
}

// Mean return mean of values
func (s *StatCalculator) Mean() float64 {
	return s.mean
}

// Min return min of values (0, if no values)
func (s *StatCalculator) Min() float64 {
	return s.min
}

// Max return max of values (0, if no values)
func (s *StatCalculator) Max() float64 {
	return s.max
}

// Sum return sum of values (0, if no values)
func (s *StatCalculator) Sum() float64 {
	return s.sum
}

// Values return values slice
func (s *StatCalculator) Values() []float64 {
	return s.values
}

// Percentile get the value which %percent% of the values are less than. This works
// just like median (where median represents the 50% point). A typical
// desire is to see the 90% point - the value that 90% of the data points
// are below, the remaining 10% are above.
//
// @param percent number representing the wished percent (between 0 and 100)
// @return the value which %percent% of the values are less than
func (s *StatCalculator) Percentile(percent float64) float64 {
	if s.count == 0 {
		return s.zero
	}
	if percent >= 100 {
		return s.Max()
	}
	percentile, err := stats.Percentile(s.values, percent)
	if err != nil {
		return math.NaN()
	}
	return percentile
}

// Count return value's count (0, if no values)
func (s *StatCalculator) Count() uint64 {
	return s.count
}

// AddAll append stat from other StatCalculator
func (s *StatCalculator) AddAll(o *StatCalculator) *StatCalculator {
	if o.Count() == 0 {
		return s
	}
	if s.count == 0 {
		s.sum = o.Sum()
	} else {
		s.sum += o.Sum()
	}
	s.count += o.Count()
	s.mergeDerivedValues(o)
	s.updateValues(o)

	return s
}

// AddValue add the value
func (s *StatCalculator) AddValue(value float64) *StatCalculator {
	if s.count == 0 {
		s.sum = value
	} else {
		s.sum += value
	}
	s.count++
	s.calculateDerivedValues(value)
	s.updateValue(value)

	return s
}

func (s *StatCalculator) updateValue(value float64) {
	s.values = append(s.values, value)
}

func (s *StatCalculator) updateValues(o *StatCalculator) {
	s.values = append(s.values, o.values...)
}

func (s *StatCalculator) calculateDerivedValues(value float64) {
	s.mean = s.sum / float64(s.count)
	if value > s.max || len(s.values) == 0 {
		s.max = value
	}
	if value < s.min || len(s.values) == 0 {
		s.min = value
	}
}

func (s *StatCalculator) mergeDerivedValues(o *StatCalculator) {
	s.mean = s.sum / float64(s.count)
	if o.Max() > s.max || len(s.values) == 0 {
		s.max = o.Max()
	}
	if o.Min() < s.min || len(s.values) == 0 {
		s.min = o.Min()
	}
}

func compareWithNaN(s float64, o float64) bool {
	if math.IsNaN(s) {
		if !math.IsNaN(o) {
			return false
		}
	} else if s != o {
		return false
	}
	return true
}

// Equal with other instance
func (s *StatCalculator) Equal(o *StatCalculator) bool {
	if s.count != o.count {
		return false
	} else {
		if !compareWithNaN(s.min, o.min) {
			return false
		}

		if !compareWithNaN(s.max, o.max) {
			return false
		}

		if !compareWithNaN(s.sum, o.sum) {
			return false
		}

		if !compareWithNaN(s.mean, o.mean) {
			return false
		}

		if s.count == 0 {
			if !compareWithNaN(s.zero, o.zero) {
				return false
			} else {
				if len(s.values) != len(o.values) {
					return false
				}
				for i := range s.values {
					if s.values[i] != o.values[i] {
						return false
					}
				}
			}
		}
	}

	return true
}
