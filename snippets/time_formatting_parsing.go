package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t)                      // 2019-04-14 08:23:18.413350843 +0800 CST m=+0.000336017
	p(t.Format(time.RFC3339)) // 2019-04-14T08:23:18+08:00

	t1, _ := time.Parse(time.RFC3339, "2019-04-14T08:23:18+08:00")
	p(t1) // 2019-04-14 08:23:18 +0800 CST

	// Format and Parse use example-based layouts
	p(t1.Format("3:04PM"))                           // 8:23AM
	p(t1.Format("Mon Jan _2 15:04:05 2006"))         // Sun Apr 14 08:23:18 2019
	p(t1.Format("2006-01-02T15:04:05.999999-07:00")) // 2019-04-14T08:23:18+08:00

	layout := "3 04 PM"
	p(time.Parse(layout, "8 41 PM")) // 0000-01-01 20:41:00 +0000 UTC <nil>

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansicLayout := "Mon Jan _2 15:04:05 2006"
	if _, e := time.Parse(ansicLayout, "8:41PM"); e != nil {
		panic(e) // panic: parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
	}
}
