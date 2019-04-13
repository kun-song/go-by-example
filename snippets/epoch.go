package main

import (
	"fmt"
	"time"
)

/**
 * 获取 Unix 时间
 */
func main() {
	p := fmt.Println

	// time -> Unix epoch
	now := time.Now()
	p(now.Unix())               // 1555171081（秒）
	p(now.UnixNano() / 1000000) // 1555171161024（微秒）
	p(now.UnixNano())           // 1555171094913124478（纳秒）

	// Unix epoch -> time
	p(time.Unix(now.Unix(), 0))     // 2019-04-14 00:03:44 +0800 CST
	p(time.Unix(0, now.UnixNano())) // 2019-04-14 00:03:44.525259431 +0800 CST
}
