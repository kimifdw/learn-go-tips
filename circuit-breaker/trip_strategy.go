package circuit_breaker

type TripStrategyFunc func(Metrics) bool

//according to consecutive fail
func ConsecutiveFailTripFunc(threshold uint64) TripStrategyFunc {
	return func(m Metrics) bool {
		return m.ConsecutiveFail >= threshold
	}
}

//according to fail
func FailTripFunc(threshold uint64) TripStrategyFunc {
	return func(m Metrics) bool {
		return m.CountFail >= threshold
	}
}

//according to fail rate
func FailRateTripFunc(rate float64, minCalls uint64) TripStrategyFunc {
	return func(m Metrics) bool {
		var currRate float64
		if m.CountAll != 0 {
			currRate = float64(m.CountFail) / float64(m.CountAll)
		}
		return m.CountAll >= minCalls && currRate >= rate
	}
}

type TripStrategyOption struct {
	Strategy                 uint
	ConsecutiveFailThreshold uint64
	FailThreshold            uint64
	FailRate                 float64
	MinCall                  uint64
}

const (
	ConsecutiveFailTrip = iota + 1
	FailTrip
	FailRateTrip
)

func ChooseTrip(op *TripStrategyOption) TripStrategyFunc {
	switch op.Strategy {
	case ConsecutiveFailTrip:
		return ConsecutiveFailTripFunc(op.ConsecutiveFailThreshold)
	case FailTrip:
		return FailTripFunc(op.FailThreshold)
	case FailRateTrip:
		fallthrough
	default:
		return FailRateTripFunc(op.FailRate, op.MinCall)
	}
}
