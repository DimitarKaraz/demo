package main

import (
	"testing"
	"bytes"
)

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := SpySleeper{}
	Countdown(buffer, 3, &spySleeper)

	got := buffer.String()
	want := "3\n2\n1\nGo!"

	if got != want || spySleeper.calls != 3 {
		t.Errorf("got %q want %q", got, want)
	}
}