package ctime

import "time"

func Sleep(t float32) {
	time.Sleep(time.Duration(t * float32(time.Second)))
}
