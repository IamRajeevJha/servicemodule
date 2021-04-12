package mygomodule

import "testing"

func TestFunc(t *testing.T) {
	var s Service
	res := s.Func1("Rajeev", 21)
	if res != "Rajeev21" {
		t.Errorf("want = %s, Got = %s", "Rajeev21", res)
	}
}
