package main

import "testing"

func TestGreeting(t *testing.T) {

	result := greeting("Code.education Rocks!")
	if result != "<b>Code.education Rocks!</b>" {
		t.Errorf("greeting() failed, expected %s, got %s",
			"<b>Code.Education Rocks!</b>", result)
	}
}