package test

import (
	"fmt"
	"testing"
)

type TestSuit struct {
	T       *testing.T
	receive interface{}
}

func (ts TestSuit) ToEqual(expected interface{}) {
	if ts.receive == expected {
		ts.T.Log("test ok")
	} else {
		ts.T.Error(fmt.Sprintf("test failed: expected(%s) receive(%s)", expected, ts.receive))
	}
}

func (ts TestSuit) Expect(receive interface{}) TestSuit {
	return TestSuit{
		T:       ts.T,
		receive: receive,
	}
}
