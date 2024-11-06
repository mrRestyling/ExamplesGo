package main

import "testing"

func TestHappy(t *testing.T) {
	text := "Egor going to Oleg, take beer and go dance with Alisa"
	expected := "Egor Oleg Alisa"
	result := Text(text)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
