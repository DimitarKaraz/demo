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

type DefaultSleeper struct {
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	Countdown(os.Stdout, 3, &DefaultSleeper{})
}

func Countdown(w io.Writer, seconds int, s Sleeper) {
	for i := seconds; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprintf(w, "Go!")
}
