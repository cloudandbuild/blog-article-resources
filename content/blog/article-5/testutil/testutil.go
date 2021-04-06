package testutil

import (
	"fmt"

	"github.com/golang/mock/gomock"
)

func NewMatcher(x interface{}, matchFn func(a, b interface{}) bool) gomock.Matcher {
	return &Matcher{
		x:       x,
		matchFn: matchFn,
	}
}

type Matcher struct {
	gomock.Matcher
	x       interface{}
	matchFn func(a, b interface{}) bool
}

func (m *Matcher) Matches(x interface{}) bool {
	return m.matchFn(m.x, x)
}
func (m *Matcher) String() string {
	return fmt.Sprintf("is equal to %v", m.x)
}
