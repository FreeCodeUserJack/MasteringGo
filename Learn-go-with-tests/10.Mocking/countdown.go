package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, "%d\n", i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout, ConfigurableSleeper{1 * time.Second, time.Sleep})
}