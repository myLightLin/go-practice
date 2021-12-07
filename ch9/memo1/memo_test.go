package memo

import (
	"testing"

	"github.com/myLightLin/go-practice/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := New(httpGetBody)
	memotest.Sequential(t, m)
}
