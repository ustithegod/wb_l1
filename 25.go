package main

import "time"

func Sleep(dur time.Duration) {
	done := make(chan struct{})

	go func() {
		<-time.After(dur)
		close(done)
	}()

	<-done
}

func Sleep2(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C
}

func main() {
	Sleep(3 * time.Second)
}
