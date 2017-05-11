package log

import "math"

func log(number Float64) float64 {
	return math.Log(number) / math.Log(2)
}
