// Driver program used to experiment with the Vector package
package main

import (
	"fmt"

	"github.com/BeautifulTovarisch/linalg/vector"
)

func main() {
	z := vector.ZeroVec([]string{"A", "B"})

	fmt.Println(z)
}
