package main

import "testing"

func TestJumpSearch_Beginning(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := 1
	result := JumpSearch(arr, s)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}

func TestJumpSearch_Middle(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := 5
	result := JumpSearch(arr, s)
	if result != 4 {
		t.Errorf("Expected 4, got %d", result)
	}
}

func TestJumpSearch_End(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := 9
	result := JumpSearch(arr, s)
	if result != 8 {
		t.Errorf("Expected 8, got %d", result)
	}
}

func TestJumpSearch_NotFound(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := 10
	result := JumpSearch(arr, s)
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}
}

func TestJumpSearch_Empty(t *testing.T) {
	arr := []int{}
	s := 1
	result := JumpSearch(arr, s)
	if result != -1 {
		t.Errorf("Expected -1, got %d", result)
	}
}

func TestJumpSearch_SingleElement(t *testing.T) {
	arr := []int{1}
	s := 1
	result := JumpSearch(arr, s)
	if result != 0 {
		t.Errorf("Expected 0, got %d", result)
	}
}
