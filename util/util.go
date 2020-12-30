/**
 * @Author: Resynz
 * @Date: 2020/12/30 13:15
 */
package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenRandom(count int64) int64 {
	return rand.Int63n(count)
}
