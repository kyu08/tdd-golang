package main

import (
	"hello/cmd/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout, &mocking.DefaultSleeper{})
}
