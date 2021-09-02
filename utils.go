package yookassa

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var isSeedAcquired = false

func UUIDGen() string {
	if !isSeedAcquired {
		rand.Seed(time.Now().UnixNano())
		isSeedAcquired = true
	}

	u := make([]byte, 16)
	rand.Read(u)

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}
