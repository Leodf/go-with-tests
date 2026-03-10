package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord = "Go!"
	countDownStart = 3
)

func Countdown(out io.Writer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}