package weight_random_choose

import (
	"github.com/golang-infrastructure/go-tuple"
	"math/rand"
)

// WeightRandomChoose 带权重的随机选择
type WeightRandomChoose[T any] struct {

	// 被带权重随机选择的目标数组
	Slice []T

	// 权重数组
	Weights     []int
	maxPoint    int
	lineSegment []int
}

func NewUseTupleSlice[T any](tupleSlice []tuple.Tuple2[T, int]) *WeightRandomChoose[T] {
	slice := make([]T, len(tupleSlice))
	weights := make([]int, len(tupleSlice))
	for index, tuple := range tupleSlice {
		slice[index] = tuple.V1
		weights[index] = tuple.V2
	}
	return New(slice, weights)
}

func New[T any](slice []T, weights []int) *WeightRandomChoose[T] {
	x := &WeightRandomChoose[T]{
		Slice: slice,
	}
	x.UpdateWeights(weights)
	return x
}

// UpdateWeights 更新权重数组，当权重有更改的时候不要直接修改Weights数组，而是通过这个方法更新权重
func (x *WeightRandomChoose[T]) UpdateWeights(weights []int) {
	lineSegment := make([]int, len(weights))
	nextPoint := 0
	for i := 0; i < len(weights); i++ {
		nextPoint += weights[i]
		lineSegment[i] = nextPoint
	}
	x.Weights = weights
	x.maxPoint = nextPoint
	x.lineSegment = lineSegment
}

// Random 带权重随机选择一个值
func (x *WeightRandomChoose[T]) Random() T {

	// 开始取随机数
	n := rand.Intn(x.maxPoint)
	index, _ := BinarySearch(x.lineSegment, n)
	return x.Slice[index]
}

func RandomChoose[T any](slice []T, weights []int) T {
	return New(slice, weights).Random()
}
