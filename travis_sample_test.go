package main

import (
	"testing"
)

func TestAddSuccess(t *testing.T) {
	const a,b = 1,1
	const ret = 2
	if x:=Add(a,b); x != ret {
		t.Errorf("Add(%v,%v) = %v, want %v",a,b,x,ret)
	}
}

func TestAddFail(t *testing.T) {
	const a,b = 1,1
	const ret = 3
	if x:=Add(a,b); x == ret {
		t.Errorf("Add(%v,%v) = %v, want %v",a,b,x,ret)
	}
}


















