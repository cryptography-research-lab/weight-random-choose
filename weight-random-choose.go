package weight_random_choose

import (
	"fmt"
	"github.com/golang-infrastructure/go-tuple"
	"math/rand"
)

// WeightRandomChoose 带权重的随机选择，创建一个struct示例多次Random性能稍微好一些
type WeightRandomChoose[T any] struct {

	// 被带权重随机选择的目标数组
	Slice []T

	// 权重数组
	Weights []int
	// 线段中最大的点是多少
	maxPoint int
	// 对线段进行分段
	lineSegment []int
}

// NewUseTupleSlice 从元组数组中创建
func NewUseTupleSlice[T any](tupleSlice []tuple.Tuple2[T, int]) (*WeightRandomChoose[T], error) {
	slice := make([]T, len(tupleSlice))
	weights := make([]int, len(tupleSlice))
	for index, t := range tupleSlice {
		slice[index] = t.V1
		weights[index] = t.V2
	}
	return New(slice, weights)
}

func New[T any](slice []T, weights []int) (*WeightRandomChoose[T], error) {
	x := &WeightRandomChoose[T]{
		Slice: slice,
	}
	return x, x.UpdateWeights(weights)
}

// UpdateWeights 更新权重数组，当权重有更改的时候不要直接修改Weights数组，而是通过这个方法更新权重
func (x *WeightRandomChoose[T]) UpdateWeights(weights []int) error {

	if err := x.validateWeights(weights); err != nil {
		return err
	}

	lineSegment := make([]int, len(weights))
	nextPoint := 0
	for i := 0; i < len(weights); i++ {
		nextPoint += weights[i]
		lineSegment[i] = nextPoint
	}
	x.Weights = weights
	x.maxPoint = nextPoint
	x.lineSegment = lineSegment

	return nil
}

// 校验权重数组是否OK
func (x *WeightRandomChoose[T]) validateWeights(weights []int) error {
	for index, v := range weights {
		if v <= 0 {
			return fmt.Errorf(fmt.Sprintf("weights index %d must greater than zero", index))
		}
	}
	return nil
}

// Random 带权重随机选择一个值
func (x *WeightRandomChoose[T]) Random() T {
	// 开始取随机数
	n := rand.Intn(x.maxPoint)
	index, _ := BinarySearch(x.lineSegment, n)
	return x.Slice[index]
}

// RandomChoose 根据权重数组随机选择一个元素
func RandomChoose[T any](slice []T, weights []int) (T, error) {
	x, err := New(slice, weights)
	if err != nil {
		var zero T
		return zero, err
	}
	return x.Random(), nil
}
