package timer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	for i := uint64(1); i < 60*60*24*4; i++ {
		d := parseTicks(i)
		// if seconds wrapped around
		if d[3] == 0 {
			// if minutes wrapped around
			if d[2] == 0 {
				// if hours wrapped around
				if d[1] == 0 {
					fmt.Println("day:", d[0])
					// if days wrapped around
					if d[0] == 100 {
						i = 0
					}
				}
			}
		}
	}
}
