package main

import "testing"

func TestFibonacci(t *testing.T){
	var v uint
	v = fibonacci(0) 
	if v != 0 { t.Error("Expected 0, got ", v) }
	v = fibonacci(1) 
	if v != 1 { t.Error("Expected 1, got ", v) }
}