// public class Die {

//     // create a random int between 1 and 6 (incl)
//     public int throw_die() {
//         int i = (int)Math.floor(6 * Math.random());
//         if (0 == i)
//             return 6;
//         else
//             return i;
//     }

// }

package main

import (
	"math/rand"
	"time"
)

func throwDie() int {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(6) + 0
	if i == 0 {
		return 6
	}
	return i
}
