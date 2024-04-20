package main

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	// Define test cases with expected outputs
	tests := []struct {
		duration time.Duration
		expected string
	}{
		{
			time.Hour + 59*time.Minute + 59*time.Second + 999*time.Millisecond + 999999*time.Nanosecond,
			"  1h 59m 59s 999ms 999999ns",
		},
		{
			59*time.Minute + 59*time.Second + 999*time.Millisecond + 999999*time.Nanosecond,
			"     59m 59s 999ms 999999ns",
		},
		{
			59*time.Second + 500*time.Millisecond + 100*time.Nanosecond,
			"         59s 500ms    100ns",
		},
		{
			500*time.Millisecond + 1*time.Nanosecond,
			"             500ms      1ns",
		},
		{
			1 * time.Nanosecond,
			"                        1ns",
		},
	}

	// Iterate through test cases, checking each one
	for _, tc := range tests {
		got := formatDuration(tc.duration)
		if got != tc.expected {
			t.Errorf("formatDuration(%v) =\n\t%q, want\n\t%q", tc.duration, got, tc.expected)
		}
	}
}
