package hcheck

import (
	"testing"
)

func TestGenCheck(t *testing.T) {
	g := GenChecker{}
	got := g.Check()
	want := Status{Pass: true, Msg: "OK"}

	if want.Pass != got.Pass || want.Msg != got.Msg {
		t.Errorf("TestGenCheck Failed : want : %#v , got : %#v ", want, got)
	}
}
