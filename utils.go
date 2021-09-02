package yookassa

import (
	"encoding/hex"
	"math/rand"
)

func UUIDGen() string {
	u := make([]byte, 16)
	rand.Read(u)

	u[8] = (u[8] | 0x80) & 0xBF
	u[6] = (u[6] | 0x40) & 0x4F

	return hex.EncodeToString(u)
}
