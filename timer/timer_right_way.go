package main

import "time"

func main() {
	c := make(chan bool)
	t := 2 * time.Second
	timer := time.NewTimer(t)

	go func() {
		time.Sleep(3 * time.Second)
	}()

	for {
		select {
		case <-timer.C:
			timer.Reset(t)
		case <-c:
			if !timer.Stop() {
				select {
				case <-timer.C: //try to drain from the channel
				default:
				}
			}
			timer.Reset(t)
		}
	}
}
