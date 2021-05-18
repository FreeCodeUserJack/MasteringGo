package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const (
	sleep = "sleep"
	write = "write"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
	
		Countdown(buffer, spySleeper)
	
		got := buffer.String()
		want := `3
2
1
Go!`
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	
		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want 4, got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep between print", func(t *testing.T) {
		want := []string{write, sleep, write, sleep, write, sleep, write}

		countdownOperationsSpy := &CountdownOperationsSpy{}
	
		Countdown(countdownOperationsSpy, countdownOperationsSpy)
	
		got := countdownOperationsSpy.Calls
	
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted calls %v, got %v", want, got)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	configurableSleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}

	configurableSleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}