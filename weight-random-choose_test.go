package weight_random_choose

import (
	"fmt"
	"testing"
)

func TestWeightRandomChoose_Random(t *testing.T) {

	slice := []rune{'A', 'B', 'C'}
	weights := []int{60, 30, 10}
	choose := New(slice, weights)

	total := 100000000
	countMap := make(map[rune]int)
	for i := 0; i < total; i++ {
		n := choose.Random()
		countMap[n] += 1
	}

	for _, char := range slice {
		count := countMap[char]
		fmt.Println(fmt.Sprintf("%s %d %f", string(char), count, float64(count)/float64(total)*float64(100)))
	}
	// Output:
	// A 60995432 60.995432
	// B 30003142 30.003142
	// C 9001426 9.001426

}
