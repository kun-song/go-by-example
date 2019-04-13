package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now) // 2019-04-13 23:37:20.760517938 +0800 CST m=+0.000348683

	// 注意提供时区！
	then := time.Date(2018, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then) // 2018-11-17 20:34:58.651387237 +0000 UTC

	p(then.Year())       // 2018
	p(then.Month())      // November
	p(then.Day())        // 17
	p(then.Hour())       // 20
	p(then.Minute())     // 34
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	p(then.Weekday()) // Saturday

	// 时间比较
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	// 时间间隔（Duration）
	diff := now.Sub(then)
	p(diff) // 3523h11m36.222392094s

	// 时间间隔以 **不同单位** 表示
	p(diff.Hours())       // 3523.220379054874
	p(diff.Minutes())     // 211393.4887861053
	p(diff.Seconds())     // 1.268364427041465e+07
	p(diff.Nanoseconds()) // 12683624826698876

	p(then.Add(diff))  // 2019-04-13 15:51:04.097003523 +0000 UTC
	p(then.Add(-diff)) // 2018-06-24 01:18:28.814583991 +0000 UTC
}
