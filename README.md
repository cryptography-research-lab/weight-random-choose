# 带权重的随机选择算法（weight random choose）

# 一、这是什么？使用场景是什么？

通常情况下当我们需要从一个数组中随机选择一个元素的时候，我们一般会在`[0, len(array)]`区间生成一个随机数，然后使用这个随机数作为数组下标从数组中取一个元素，在这种情况对于数组中的每个元素被选中的概率都是`1/n`，但是如果我们需要数组中的每个元素被选中的概率都不同呢，比如有一个数组：

```json
["写代码", "看书", "打游戏"]
```

我们期望数组中的每个元素被选中的概率都不同，比如对于“写代码”的权重高一些，对于其它事情的权重低一些，权重数组大概是这样：

```json
[2, 1, 1]
```

这个时候对于每件事情随机选中的概率期望为：

```
写代码 --> 50%
看书 --> 25%
打游戏 --> 25%
```

这个仓库的目标就是如何实现这种带权重的随机分布。

# 二、安装

```bash
go get -u github.com/cryptography-research-lab/weight-random-choose 
```

# 三、API示例

```go
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
```

# 四、算法详解

TODO 





















