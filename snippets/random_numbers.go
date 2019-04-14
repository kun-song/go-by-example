package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * `math/rand` 伪随机数生成
 */

func main() {
	p := fmt.Println

	// [0, 100) 前开后闭
	p(rand.Intn(100), ",", rand.Intn(100)) // 81 , 87，重复运行结果相同

	// [0, 1.0)
	p(rand.Float64()) // 0.6645600532184904

	// [5.0, 10.0)
	p(rand.Float64()*5 + 5) // 7.1885709359349015

	/**
	 * 1. 前面的随机数，每次运行结果都相同；若要不同，需提供 seed
	 * 2. `math/rand` 不能用于加密场景，加密可用 `crypto/rand`
	 */
	seed := rand.NewSource(time.Now().UnixNano())
	rand2 := rand.New(seed)
	p(rand2.Intn(100), ",", rand2.Intn(100)) // 21 , 61，重复运行结果不同

	// 若两个 rand 的 seed 相同，则生成的随机数也相同
	s1 := rand.NewSource(666)
	r1 := rand.New(s1)
	p(r1.Intn(100), ",", r1.Intn(100)) // 12 , 63

	s2 := rand.NewSource(666)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",", r2.Intn(100)) // 12 , 63
}
