package main

import (
	"fmt"
	weight_random_choose "github.com/cryptography-research-lab/weight-random-choose"
)

func main() {

	// slice是要被选择的数组
	slice := []rune{'A', 'B', 'C'}
	// weights是与slice一一对应的每个元素被选中的的权重
	weights := []int{60, 30, 10}

	// 当需要运行多次带权重的随机选择，创建一个struct性能会比较好
	choose := weight_random_choose.New(slice, weights)

	// 运行足够大的次数，统计每个字符真正被选中的次数
	total := 100000000
	countMap := make(map[rune]int)
	for i := 0; i < total; i++ {
		n := choose.Random()
		countMap[n] += 1
	}

	// 看一下这些字符真正被选中的次数是否符合刚开始的期望
	for _, char := range slice {
		count := countMap[char]
		fmt.Println(fmt.Sprintf("%s %d %f", string(char), count, float64(count)/float64(total)*float64(100)))
	}
	// Output:
	// A 60994473 60.994473
	// B 30007986 30.007986
	// C 8997541 8.997541

}
