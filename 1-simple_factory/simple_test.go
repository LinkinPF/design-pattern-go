package __simple_factory

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	fmt.Printf("api's type is %T\n", api)
	s := api.Say("zcy")
	if s != "Hi, zcy" {
		t.Fatal("TestType1 test failed")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("zcy")
	if s != "Hi, zcy" {
		t.Fatal("TestType2 test failed")
	}
}
