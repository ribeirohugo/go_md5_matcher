package fault

import (
	"time"
)

type Timer interface {
	Now() time.Time
}

type RealTimer struct{}

func (r RealTimer) Now() int64 {
	return time.Now().Unix()
}

type MockTimer struct {
	expectedTime int64
}

func (m MockTimer) Now() int64 {
	return m.expectedTime
}
