package popcount

import (
	"strconv"
)

func PopCount4(x uint64) int {
	var count int
	p := strconv.FormatUint(x, 2)
	for i := range p {
		count += int(p[i]) & 1
	}
	return count
}
