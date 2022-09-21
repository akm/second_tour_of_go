package main

import "time"

var NowFunc = time.Now

func Now() time.Time {
	return NowFunc()
}
