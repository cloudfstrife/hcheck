package hcheck

import (
	"testing"
)

//TestChecker checker
type TestChecker struct {
}

//Check Testing
func (dc *TestChecker) Check() Status {
	//do  message queue connection check
	return Status{
		Pass: false,
		Msg:  "Testing",
	}
}

func TestRegister(t *testing.T) {
	want := &TestChecker{}
	name := "Testing"
	Register(name, want)
	defer Unregister(name)
	got := checkers[name]

	if want != got {
		t.Errorf("TestRegister Failed : want : %#v , got : %#v ", want, got)
	}
}

func TestUnregister(t *testing.T) {
	name := "Testing"
	Register(name, &TestChecker{})
	Unregister(name)
	got := checkers[name]

	if got != nil {
		t.Errorf("TestRegister Failed : want : %#v , got : %#v ", nil, got)
	}
}

func TestCheck(t *testing.T) {
	genKey := "Gen"
	testKey := "Testing"
	genWant := Status{Pass: true, Msg: "OK"}
	testWant := Status{Pass: false, Msg: "Testing"}

	Register(testKey, &TestChecker{})
	defer Unregister(testKey)
	got := Check()

	if len(got) != 2 {
		t.Errorf("TestCheck Failed : got : %#v ", got)
	}
	genGot, ok := got[genKey]
	if !ok || genWant.Pass != genGot.Pass || genWant.Msg != genGot.Msg {
		t.Errorf("TestGenCheck Failed : want : %#v , got : %#v ", genWant, genGot)
	}
	testGot, ok := got[testKey]
	if !ok || testWant.Pass != testGot.Pass || testWant.Msg != testGot.Msg {
		t.Errorf("TestGenCheck Failed : want : %#v , got : %#v ", testWant, testGot)
	}
}
